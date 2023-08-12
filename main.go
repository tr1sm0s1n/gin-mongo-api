package main

import (
	"context"
	"log"
	"net/http"

	"github.com/DEMYSTIF/go-mongo-api/config"
	"github.com/DEMYSTIF/go-mongo-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func create(c *gin.Context, client *mongo.Client) {
	var newCertificate models.Certificate
	if err := c.ShouldBindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	coll := config.GetCollection(client, "certificates")
	result, err := coll.InsertOne(context.TODO(), newCertificate)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Duplicate Key"})
		return
		}
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"_id": result.InsertedID})
}

func readAll(c *gin.Context, client *mongo.Client) {
	coll := config.GetCollection(client, "certificates")
	result, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	var results []models.Certificate
	if err = result.All(context.TODO(), &results); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func main() {
	client, err := config.ConnectClient()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	router.POST("/create", func(ctx *gin.Context) {
		create(ctx, client)
	})
	router.GET("/read", func(ctx *gin.Context) {
		readAll(ctx, client)
	})
	router.Run("localhost:8080")
}
