package models

import (
	"activity/database"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Graph = database.Database.Collection("graphs")
)

type GraphProps struct {
	URL           string    `bson:"url"`
	ReferenceDate time.Time `bson:"referenceDate"`
	CreatedAt     time.Time `bson:"createdAt"`
	UpdatedAt     time.Time `bson:"updatedAt"`
}

func (g *GraphProps) MarshalBSON() ([]byte, error) {
	if g.CreatedAt.IsZero() {
		g.CreatedAt = time.Now()
	}
	g.UpdatedAt = time.Now()

	type my GraphProps
	return bson.Marshal((*my)(g))
}

func GetLastGraph() (res GraphProps) {
	Graph.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{
		"createdAt": -1,
	})).Decode(&res)
	return res
}

func AddNewGraph(graphURL string, referenceDate time.Time) {
	Graph.InsertOne(context.Background(), &GraphProps{
		URL:           graphURL,
		ReferenceDate: referenceDate,
	})
}
