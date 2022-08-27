package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"function/config"
	"net/http"

	"github.com/diamondburned/arikawa/v3/discord"
)

func SendAllGuilds(message *discord.Message) {
	data := struct {
		ChannelIDString string          `json:"channelId" binding:"required"`
		MessageIDString string          `json:"messageId" binding:"required"`
		Embeds          []discord.Embed `json:"embeds" binding:"required"`
	}{
		ChannelIDString: message.ChannelID.String(),
		MessageIDString: message.ID.String(),
		Embeds:          message.Embeds,
	}
	dataBytes, _ := json.Marshal(data)
	reader := bytes.NewReader(dataBytes)
	http.Post(fmt.Sprintf("%s/send-all", config.ActivityBaseURL), "application/json", reader)
}
