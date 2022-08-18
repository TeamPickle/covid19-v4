package send

import (
	"context"
	"fmt"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func makeModalResponse(selectedType string) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.ModalResponse,
		Data: &api.InteractionResponseData{
			Title:    option.NewNullableString(fmt.Sprintf("전송할 내용을 입력해주세요. (%s)", sendType(selectedType).Name())),
			CustomID: option.NewNullableString(selectedType),
			Components: &discord.ContainerComponents{
				&discord.ActionRowComponent{
					&discord.TextInputComponent{
						Style:    discord.TextInputParagraphStyle,
						Required: true,
						CustomID: discord.ComponentID(selectedType),
						Label:    "내용",
					},
				},
			},
		},
	}
}

func (*SendCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	selectedType := data.Options.Find("type").String()
	return makeModalResponse(selectedType)
}
