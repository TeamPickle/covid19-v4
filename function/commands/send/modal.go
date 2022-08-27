package send

import (
	"context"
	"fmt"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (*SendModal) Handle(ctx context.Context, data *discord.ModalInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	selectedType := sendType(data.CustomID)
	actionRowComponent := *data.Components[0].(*discord.ActionRowComponent)
	textComponent := actionRowComponent[0].(*discord.TextInputComponent)
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString(fmt.Sprintf("이렇게 __**%s**__ 메시지를 전송하시겠습니까?", selectedType)),
		Embeds:  &[]discord.Embed{*makeEmbedWithContent(selectedType, textComponent.Value.Val)},
		Components: &discord.ContainerComponents{
			&discord.ActionRowComponent{
				&discord.ButtonComponent{
					Style:    discord.PrimaryButtonStyle(),
					CustomID: "send",
					Label:    "전송",
				},
				&discord.ButtonComponent{
					Style:    discord.DangerButtonStyle(),
					CustomID: "cancel",
					Label:    "취소",
				},
				&discord.ButtonComponent{
					Style:    discord.SecondaryButtonStyle(),
					CustomID: "typo-" + data.CustomID,
					Label:    "다시 작성",
				},
			},
		},
	})
}
