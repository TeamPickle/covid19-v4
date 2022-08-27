package main

import (
	"context"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func handler(ctx context.Context, interaction discord.InteractionEvent) *api.InteractionResponse {
	if interaction.Data.InteractionType() == discord.PingInteractionType {
		return &api.InteractionResponse{Type: api.PongInteraction}
	}

	return func() (response *api.InteractionResponse) {
		defer func() {
			if r := recover(); r != nil {
				response = utils.MakeErrorMessage(interaction, r)
			}
		}()

		switch i := interaction.Data.(type) {
		case *discord.CommandInteraction:
			return commandHandler.Handle(ctx, i, interaction)
		case *discord.AutocompleteInteraction:
			return autoCompleteHandler.Handle(ctx, i)
		case discord.ComponentInteraction:
			return componentHandler.Handle(ctx, i, interaction)
		case *discord.ModalInteraction:
			return modalHandler.Handle(ctx, i, interaction)
		}
		return
	}()
}
