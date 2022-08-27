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
	embedsBytes, _ := json.Marshal(message.Embeds)
	reader := bytes.NewReader(embedsBytes)
	http.Post(fmt.Sprintf("%s/send-all", config.ActivityBaseURL), "application/json", reader)
}
