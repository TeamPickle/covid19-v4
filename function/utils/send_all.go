package utils

import (
	"fmt"
	"function/config"
	"net/http"

	"github.com/diamondburned/arikawa/v3/discord"
)

func SendAllGuilds(message *discord.Message) {
	http.Post(fmt.Sprintf("%s/channel/%s/message/%s", config.ActivityBaseURL, message.ChannelID, message.ID), "", nil)
}
