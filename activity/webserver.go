package main

import (
	"activity/config"
	"fmt"
	"net/http"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/gin-gonic/gin"
)

func runWebServer(m *shard.Manager) {
	client := api.NewClient("Bot " + config.Token)
	r := gin.Default()

	r.POST("/send-all", func(c *gin.Context) {
		data := struct {
			ChannelIDString string          `json:"channelId" binding:"required"`
			MessageIDString string          `json:"messageId" binding:"required"`
			Embeds          []discord.Embed `json:"embeds" binding:"required"`
		}{}
		c.Bind(&data)

		channelID := discord.NullChannelID
		messageID := discord.NullMessageID

		if c, err := discord.ParseSnowflake(data.ChannelIDString); err == nil {
			channelID = discord.ChannelID(c)
		}
		if m, err := discord.ParseSnowflake(data.MessageIDString); err == nil {
			messageID = discord.MessageID(m)
		}

		editMessage, _ := client.SendTextReply(channelID, "전송 시작", messageID)
		editMessageId := editMessage.ID

		SendAllServers(m, data.Embeds[0], func(curr, all, errors int) {
			if curr == all {
				client.EditMessage(channelID, editMessageId, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors))
				return
			}
			if curr%50 == 0 {
				client.EditMessage(channelID, editMessageId, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors))
			}
		})
		c.JSON(http.StatusOK, gin.H{})
	})
	r.Run("0.0.0.0:" + config.Port)
}
