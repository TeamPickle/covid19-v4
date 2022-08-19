package main

import (
	"context"
	"fmt"
	"function/config"
	"log"
	"net/http"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/gin-gonic/gin"
)

func main() {
	s := session.New("Bot " + config.Token)
	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	defer s.Close()

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
		message, _ := s.Message(channelID, messageID)
		SendAllServers(s, message.Embeds[0], func(curr, all, errors int) {
			if curr%10 == 0 {
				s.EditMessage(channelID, messageID, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors), message.Embeds...)
			}
			if curr == all {
				s.EditMessage(channelID, messageID, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors), message.Embeds...)
			}
		})
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run()
}
