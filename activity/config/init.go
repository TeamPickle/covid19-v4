package config

import (
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joho/godotenv"
)

var (
	Port            string
	Token           string
	LogWebhookID    discord.WebhookID
	LogWebhookToken string
	MongoDBURL      string
)

func init() {
	godotenv.Load()
	if Port = os.Getenv("PORT"); Port == "" {
		Port = "3000"
	}
	if Token = os.Getenv("BOT_TOKEN"); Token == "" {
		panic("BOT_TOKEN is not set")
	}
	if channelID, err := discord.ParseSnowflake(os.Getenv("LOG_WEBHOOK_ID")); err == nil {
		LogWebhookID = discord.WebhookID(channelID)
	}
	LogWebhookToken = os.Getenv("LOG_WEBHOOK_TOKEN")
	if MongoDBURL = os.Getenv("MONGO_DB_URL"); MongoDBURL == "" {
		panic("MONGO_DB_URL is not set")
	}
}
