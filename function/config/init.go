package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Token      string
	OpenAPIkey string
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	Token = os.Getenv("BOT_TOKEN")
	OpenAPIkey = os.Getenv("OPEN_API_KEY")
}
