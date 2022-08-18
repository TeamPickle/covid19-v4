package world

import (
	"context"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (c *WorldCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	embedData, totalPages := getMainData(ctx)
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds:     &[]discord.Embed{makeMainEmbed(&embedData)},
		Components: makeComponents(1, totalPages),
	})
}
