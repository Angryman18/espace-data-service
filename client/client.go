package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Db *mongo.Database
	// Redis *mongo.Redis
}

func NewClient(db *mongo.Database) *Client {
	return &Client{Db: db}
}

func ConnectMongo(mongoURI string) (*mongo.Database, *mongo.Client) {
	opt := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		panic("Mongo Connection Error")
	}
	database := mongoClient.Database("espace")
	return database, mongoClient
}
