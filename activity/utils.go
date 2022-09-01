package main

import (
	"activity/models"
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
)

func SendAllServers(m *shard.Manager, embed discord.Embed, onProgress func(curr int, all int, errors int)) {
	guard := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	curr := 0
	errorCount := 0
	allGuilds := 0

	m.ForEach(func(shard shard.Shard) {
		state := shard.(*state.State)
		guilds, err := state.Guilds()
		if err != nil {
			panic(err)
		}
		allGuilds += len(guilds)

		go func() {
			for _, guild := range guilds {
				guard <- struct{}{}
				wg.Add(1)
				go func(guild discord.Guild) {
					defer func() {
						<-guard
						wg.Done()
						curr++
						onProgress(curr, allGuilds, errorCount)
					}()
					channelID := models.GetChannelIDSettingSnowflake(context.TODO(), guild.ID.String())
					channelEmbed := embed

					if channelID != discord.NullChannelID {
						if _, err := state.SendEmbeds(channelID, channelEmbed); err == nil {
							return
						}
						channelEmbed.Footer = &discord.EmbedFooter{Text: "설정 되어있는 채널의 접근 권한이 없어, 메시지 전송 채널이 변경 되었습니다."}
					} else {
						channelEmbed.Footer = &discord.EmbedFooter{Text: "채널이 설정 되어있지 않습니다. 앞으로 해당 채널에 메시지를 전송합니다."}
					}

					channels, _ := state.Channels(guild.ID)
					sort.Slice(channels, func(i, j int) bool {
						return channels[i].Position < channels[j].Position
					})
					err := errors.New("no channels")
					for _, channel := range channels {
						if _, err = state.SendEmbeds(channel.ID, channelEmbed); err == nil {
							models.UpdateChannelIDSetting(context.TODO(), guild.ID.String(), channel.ID.String())
							return
						}
					}

					// TODO: 이래도 안 되면 추가적인 처리가 필요합니다.
					errorCount++
					fmt.Println(err, guild.ID, guild.Name, guild.OwnerID)
				}(guild)
			}
		}()
	})

	wg.Wait()
	onProgress(allGuilds, allGuilds, errorCount)
}
