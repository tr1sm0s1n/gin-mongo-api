package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/DEMYSTIF/go-mongo-api/config"
	"github.com/DEMYSTIF/go-mongo-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOne(c *gin.Context, client *mongo.Client) {
	var newCertificate models.Certificate
	if err := c.ShouldBindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	coll := config.GetCollection(client, "certificates")
	if _, err := coll.InsertOne(context.TODO(), newCertificate); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Duplicate Key"})
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func ReadAll(c *gin.Context, client *mongo.Client) {
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

func ReadOne(c *gin.Context, client *mongo.Client) {
	var oldCertificate models.Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	coll := config.GetCollection(client, "certificates")
	filter := bson.D{{Key: "_id", Value: id}}
	if err = coll.FindOne(context.TODO(), filter).Decode(&oldCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, oldCertificate)
}

func UpdateOne(c *gin.Context, client *mongo.Client) {
	var oldCertificate models.Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err = c.ShouldBindJSON(&oldCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	coll := config.GetCollection(client, "certificates")
	filter := bson.D{{Key: "_id", Value: id}}
	opts := options.Replace().SetUpsert(true)
	if _, err := coll.ReplaceOne(context.TODO(), filter, oldCertificate, opts); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, oldCertificate)
}

func DeleteOne(c *gin.Context, client *mongo.Client) {
	var oldCertificate models.Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	coll := config.GetCollection(client, "certificates")
	filter := bson.D{{Key: "_id", Value: id}}
	if err = coll.FindOne(context.TODO(), filter).Decode(&oldCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	
	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, oldCertificate)
}
