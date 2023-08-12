package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	return client, err
}

func GetCollection(client *mongo.Client, name string) *mongo.Collection {
	collection := client.Database("gin-mongo").Collection(name)
	return collection
}
