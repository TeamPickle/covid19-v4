package main

import (
	"context"
	"function/models"
	"sort"
	"sync"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/session"
)

func SendAllServers(session *session.Session, embed discord.Embed, onProgress func(curr int, all int, errors int)) {
	guilds, err := session.Guilds(0)
	if err != nil {
		panic(err)
	}

	guard := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	curr := 0
	errors := 0

	for _, guild := range guilds {
		guard <- struct{}{}
		wg.Add(1)
		go func(guild discord.Guild) {
			defer func() {
				<-guard
				wg.Done()
				curr++
				onProgress(curr, len(guilds), errors)
			}()
			channelID := models.GetChannelIDSettingSnowflake(context.TODO(), guild.ID.String())
			channelEmbed := embed

			if channelID != discord.NullChannelID {
				if _, err := session.SendEmbeds(channelID, channelEmbed); err == nil {
					return
				}
				channelEmbed.Footer = &discord.EmbedFooter{Text: "설정 되어있는 채널의 접근 권한이 없어, 메시지 전송 채널이 변경 되었습니다."}
			} else {
				channelEmbed.Footer = &discord.EmbedFooter{Text: "채널이 설정 되어있지 않습니다. 앞으로 해당 채널에 메시지를 전송합니다."}
			}

			channels, _ := session.Channels(guild.ID)
			sort.Slice(channels, func(i, j int) bool {
				return channels[i].Position < channels[j].Position
			})
			for _, channel := range channels {
				if _, err := session.SendEmbeds(channel.ID, channelEmbed); err == nil {
					models.UpdateChannelIDSetting(context.TODO(), guild.ID.String(), channel.ID.String())
					return
				}
			}

			// TODO: 이래도 안 되면 추가적인 처리가 필요합니다.
			errors++
		}(guild)
	}
	wg.Wait()
	onProgress(len(guilds), len(guilds), errors)
}
