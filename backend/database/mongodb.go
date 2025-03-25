package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"request-debug/config"
)

func NewMongoDB() *mongo.Client {
	client, err := mongo.Connect(options.Client().SetAppName(config.Conf.App.Name).ApplyURI(config.Conf.Database.Uri))
	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	return client
}
