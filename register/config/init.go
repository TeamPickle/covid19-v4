package config

import (
	"log"
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

	if testGuildId, err := discord.ParseSnowflake(os.Getenv("TEST_GUILD_ID")); err != nil {
		log.Println("TEST_GUILD_ID is not valid.")
	} else {
		TestGuildId = discord.GuildID(testGuildId)
	}

	if appID, err := discord.ParseSnowflake(os.Getenv("APP_ID")); err != nil {
		panic("APP_ID is not set or invalid")
	} else {
		AppID = discord.AppID(appID)
	}
}
