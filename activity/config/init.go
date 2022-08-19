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
}
