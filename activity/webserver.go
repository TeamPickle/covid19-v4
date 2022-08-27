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
	me, _ := client.Me()
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
		editMessageId := messageID
		if message.Author.ID != me.ID {
			m, _ := client.SendTextReply(channelID, "전송 시작", messageID)
			editMessageId = m.ID
		}
		SendAllServers(m, message.Embeds[0], func(curr, all, errors int) {
			fmt.Println(curr)
			if curr%10 == 0 {
				client.EditMessage(channelID, editMessageId, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors), message.Embeds...)
			}
			if curr == all {
				client.EditMessage(channelID, editMessageId, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors), message.Embeds...)
			}
		})
		c.JSON(http.StatusOK, gin.H{})
	})
	r.Run("0.0.0.0:" + config.Port)
}
