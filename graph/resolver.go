package graph

import (
	"github.com/Angryman18/espace-data-service/client"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *client.Client
}

type DB struct {
	*mongo.Database
}

func NewDB(client *mongo.Database) *DB {
	return &DB{Database: client}
}
