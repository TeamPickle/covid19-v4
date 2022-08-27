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

		editMessageId := messageID
		message, _ := client.Message(channelID, messageID)

		if message.Author.ID != me.ID {
			m, _ := client.SendTextReply(channelID, "전송 시작", messageID)
			editMessageId = m.ID
		}
		SendAllServers(m, data.Embeds[0], func(curr, all, errors int) {
			fmt.Println(curr)
			if curr == all {
				client.EditMessage(channelID, editMessageId, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors), data.Embeds...)
				return
			}
			if curr%20 == 0 {
				client.EditMessage(channelID, editMessageId, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors), data.Embeds...)
			}
		})
		c.JSON(http.StatusOK, gin.H{})
	})
	r.Run("0.0.0.0:" + config.Port)
}
