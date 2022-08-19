package database

import (
	"context"
	"function/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

const (
	databaseName = "covid19"
)

func init() {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var err error
	if Client, err = mongo.Connect(timeout, options.Client().ApplyURI(config.MongoDBURL)); err != nil {
		panic(err)
	}

	if err = Client.Ping(timeout, nil); err != nil {
		panic(err)
	}

	Database = Client.Database(databaseName)
}
