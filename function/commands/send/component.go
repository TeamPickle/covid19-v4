package send

import (
	"context"
	"function/utils"
	"strings"
	"time"

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
	go utils.SendAllGuilds(rawRequest.Message)
	time.Sleep(1 * time.Second)
	return &api.InteractionResponse{
		Type: api.UpdateMessage,
		Data: &api.InteractionResponseData{
			Content:    option.NewNullableString("전송중..."),
			Embeds:     &rawRequest.Message.Embeds,
			Components: &discord.ContainerComponents{},
		},
	}
}
