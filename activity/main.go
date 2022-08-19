package main

import (
	"activity/config"
	"activity/cron"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/gin-gonic/gin"
)

func main() {
	newShard := state.NewShardFunc(func(m *shard.Manager, s *state.State) {
		s.AddIntents(gateway.IntentGuilds)
	})
	m, err := shard.NewManager("Bot "+config.Token, newShard)
	s := m.Shard(0).(*state.State)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	defer m.Close()

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
		message, _ := m.Shard(0).(*state.State).Message(channelID, messageID)
		SendAllServers(m, message.Embeds[0], func(curr, all, errors int) {
			if curr%10 == 0 {
				s.EditMessage(channelID, messageID, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors), message.Embeds...)
			}
			if curr == all {
				s.EditMessage(channelID, messageID, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors), message.Embeds...)
			}
		})
		c.JSON(http.StatusOK, gin.H{})
	})

	go cron.Start(m)
	r.Run()
}
