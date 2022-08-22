package config

import (
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joho/godotenv"
)

var (
	OpenAPIkey      string
	LogWebhookID    discord.WebhookID = discord.NullWebhookID
	LogWebhookToken string
	MongoDBURL      string
	PublicKey       string
	ActivityBaseURL string
)

func init() {
	godotenv.Load()

	OpenAPIkey = os.Getenv("OPEN_API_KEY")
	if channelID, err := discord.ParseSnowflake(os.Getenv("LOG_WEBHOOK_ID")); err == nil {
		LogWebhookID = discord.WebhookID(channelID)
	}
	LogWebhookToken = os.Getenv("LOG_WEBHOOK_TOKEN")
	if MongoDBURL = os.Getenv("MONGO_DB_URL"); MongoDBURL == "" {
		panic("MONGO_DB_URL is not set")
	}
	PublicKey = os.Getenv("VALIDATE_PUBLIC_KEY")
	ActivityBaseURL = os.Getenv("ACTIVITY_BASE_URL")
}
