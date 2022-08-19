package send

import (
	"context"
	"fmt"
	"function/config"
	"function/utils"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (*SendComponent) Handle(ctx context.Context, data discord.ComponentInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	idString := string(data.ID())
	if idString == "cancel" {
		return &api.InteractionResponse{
			Type: api.UpdateMessage,
			Data: &api.InteractionResponseData{
				Content:    option.NewNullableString("취소되었습니다."),
				Embeds:     &rawRequest.Message.Embeds,
				Components: &discord.ContainerComponents{},
			},
		}
	}
	if strings.HasPrefix(idString, "typo") {
		return makeModalResponse(idString[5:])
	}
	go utils.SendAllServers(rawRequest.Message.Embeds[0], func(curr, all, errors int) {
		client := api.NewClient("Bot " + config.Token)
		if curr%10 == 0 {
			client.EditMessage(rawRequest.Message.ChannelID, rawRequest.Message.ID, fmt.Sprintf("전송 중... %d/%d, errors: %d", curr, all, errors), rawRequest.Message.Embeds...)
		}
		if curr == all {
			client.EditMessage(rawRequest.Message.ChannelID, rawRequest.Message.ID, fmt.Sprintf("전송 완료 to %d servers with %d errors", all, errors), rawRequest.Message.Embeds...)
		}
	})
	return &api.InteractionResponse{
		Type: api.UpdateMessage,
		Data: &api.InteractionResponseData{
			Content:    option.NewNullableString("전송중..."),
			Embeds:     &rawRequest.Message.Embeds,
			Components: &discord.ContainerComponents{},
		},
	}
}
