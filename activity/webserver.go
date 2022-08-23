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

	r.POST("/channel/:channelId/message/:messageId", func(c *gin.Context) {
		channelID := discord.NullChannelID
		messageID := discord.NullMessageID
		if c, err := discord.ParseSnowflake(c.Param("channelId")); err == nil {
			channelID = discord.ChannelID(c)
		}
		if c, err := discord.ParseSnowflake(c.Param("messageId")); err == nil {
			messageID = discord.MessageID(c)
		}
		message, _ := client.Message(channelID, messageID)
		SendAllServers(m, message.Embeds[0], func(curr, all, errors int) {
			if curr%10 == 0 {
				client.EditMessage(channelID, messageID, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors), message.Embeds...)
			}
			if curr == all {
				client.EditMessage(channelID, messageID, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors), message.Embeds...)
			}
		})
		c.JSON(http.StatusOK, gin.H{})
	})
	r.Run("0.0.0.0:" + config.Port)
}
