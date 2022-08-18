package database

import (
	"context"
	"function/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database *mongo.Database
)

const (
	databaseName = "covid19"
)

func Connect() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoDBURL))
	if err != nil {
		panic(err)
	}

	database = client.Database(databaseName)
	Location = database.Collection(locationCollectionName)
}
