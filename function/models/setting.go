package models

import (
	"function/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	Setting = database.Database.Collection("settings")
)

type SettingUpdateProps struct {
	Channel   string    `bson:"channel"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type ChannelProps struct {
	SettingUpdateProps
	ID string `bson:"_id"`
}

func (s *SettingUpdateProps) MarshalBSON() ([]byte, error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	s.UpdatedAt = time.Now()

	type my SettingUpdateProps
	return bson.Marshal((*my)(s))
}
