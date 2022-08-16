package status

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func handleRegion() *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content: option.NewNullableString("지역 현황"),
		},
	}
}

func handleDomestic() *api.InteractionResponse {
	_, err := parseNCov()
	if err != nil {
		return &api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Content: option.NewNullableString("오류가 발생했습니다."),
			},
		}
	}
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content: option.NewNullableString("국내 현황"),
		},
	}
}

func (c *StatusCommand) Handle(ctx context.Context, interaction *discord.CommandInteraction) *api.InteractionResponse {
	if len(interaction.Options) == 0 {
		return handleDomestic()
	}
	return handleRegion()
}
