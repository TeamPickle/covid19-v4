package channel

import (
	"context"
	"fmt"
	"function/models"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func getOptions(data *discord.CommandInteraction, rawRequest discord.InteractionEvent) (channelIDString, guildIDString string) {
	channel := data.Options.Find("channel")
	guildIDString = rawRequest.GuildID.String()
	channelIDString = channel.String()
	return
}

func (*ChannelCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	channelIDString, guildIDString := getOptions(data, rawRequest)
	if channelIDString == "" {
		channelIDString := models.GetChannelIDSetting(ctx, guildIDString)
		return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
			Content: option.NewNullableString(fmt.Sprintf("현재 공지를 보내는 채널은 <#%s> 입니다.", channelIDString)),
			Flags:   api.EphemeralResponse,
		})
	}
	models.UpdateChannelIDSetting(ctx, guildIDString, channelIDString)
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString(fmt.Sprintf("<#%s>(이)가 전체공지를 띄울 채널로 설정되었습니다.", channelIDString)),
		Flags:   api.EphemeralResponse,
	})
}
