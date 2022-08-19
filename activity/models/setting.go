package models

import (
	"activity/database"
	"context"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetChannelIDSetting(ctx context.Context, guildIDString string) (channelIDString string) {
	settingUpdate := SettingUpdateProps{}
	Setting.FindOne(ctx, bson.M{"_id": guildIDString}).Decode(&settingUpdate)
	return settingUpdate.Channel
}

func GetChannelIDSettingSnowflake(ctx context.Context, guildIDString string) (channelID discord.ChannelID) {
	snowflake, err := discord.ParseSnowflake(GetChannelIDSetting(ctx, guildIDString))
	if err != nil {
		return discord.NullChannelID
	}
	return discord.ChannelID(snowflake)
}

func UpdateChannelIDSetting(ctx context.Context, guildIDString, channelIDString string) {
	settingUpdate := SettingUpdateProps{}
	Setting.FindOne(ctx, bson.M{"_id": guildIDString}).Decode(&settingUpdate)
	settingUpdate.Channel = channelIDString
	_, err := Setting.ReplaceOne(ctx, bson.M{"_id": guildIDString}, &settingUpdate, options.Replace().SetUpsert(true))
	if err != nil {
		panic(err)
	}
}
