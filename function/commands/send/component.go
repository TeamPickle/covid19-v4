package send

import (
	"context"
	"fmt"
	"function/config"
	"net/http"
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
	go http.Post(fmt.Sprintf("%s/channel/%s/message/%s", config.ActivityBaseURL, rawRequest.Message.ChannelID, rawRequest.Message.ID), "", nil)
	return &api.InteractionResponse{
		Type: api.UpdateMessage,
		Data: &api.InteractionResponseData{
			Content:    option.NewNullableString("전송중..."),
			Embeds:     &rawRequest.Message.Embeds,
			Components: &discord.ContainerComponents{},
		},
	}
}
