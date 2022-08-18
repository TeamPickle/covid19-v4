package models

import (
	"function/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	Location = database.Database.Collection("locations")
)

type LocationUpdateProps struct {
	Location  string    `bson:"location"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type LocationProps struct {
	LocationUpdateProps
	ID string `bson:"_id"`
}

func (l *LocationUpdateProps) MarshalBSON() ([]byte, error) {
	if l.CreatedAt.IsZero() {
		l.CreatedAt = time.Now()
	}
	l.UpdatedAt = time.Now()

	type my LocationUpdateProps
	return bson.Marshal((*my)(l))
}
