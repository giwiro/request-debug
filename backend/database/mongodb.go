package database

import (
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"request-debug/config"
)

func NewMongoDB() *mongo.Client {
	bsonOpts := &options.BSONOptions{
		ObjectIDAsHexString: true,
	}
	client, err := mongo.Connect(
		options.Client().
			SetAppName(config.Conf.App.Name).
			ApplyURI(config.Conf.Database.Uri).
			SetBSONOptions(bsonOpts),
	)
	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	return client
}

func GetCollection(db *mongo.Client, name string) *mongo.Collection {
	return db.Database(config.Conf.Database.DBName).Collection(name)
}

func GetStringId(mongoId interface{}) (string, error) {
	objectID, ok := mongoId.(bson.ObjectID)
	if !ok {
		return "", errors.New("invalid ObjectID")
	}

	return objectID.Hex(), nil
}
