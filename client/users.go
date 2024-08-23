package client

import (
	"context"
	"log"

	"github.com/Angryman18/espace-data-service/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) GetUsers() []*model.UserData {
	cursor, err := c.Db.Collection("users").Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Error Occured")
	}
	var data []*model.UserData
	cursor.All(context.TODO(), &data)
	return data
}
