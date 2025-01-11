package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tr1sm0s1n/gin-mongo-api/config"
	"github.com/tr1sm0s1n/gin-mongo-api/controllers"
	"github.com/tr1sm0s1n/gin-mongo-api/middlewares"
)

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
	router.Use(middlewares.Authority())
	router.POST("/create", func(ctx *gin.Context) {
		controllers.CreateOne(ctx, client)
	})
	router.GET("/read", func(ctx *gin.Context) {
		controllers.ReadAll(ctx, client)
	})
	router.GET("/read/:id", func(ctx *gin.Context) {
		controllers.ReadOne(ctx, client)
	})
	router.PUT("/update/:id", func(ctx *gin.Context) {
		controllers.UpdateOne(ctx, client)
	})
	router.DELETE("/delete/:id", func(ctx *gin.Context) {
		controllers.DeleteOne(ctx, client)
	})
	router.Run("localhost:8080")
}
