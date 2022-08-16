package config

import (
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joho/godotenv"
)

var (
	Token       string
	TestGuildId discord.GuildID = discord.NullGuildID
	AppID       discord.AppID   = discord.NullAppID
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	Token = os.Getenv("BOT_TOKEN")
}
