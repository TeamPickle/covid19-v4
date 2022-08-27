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

		if interaction.Data.InteractionType() == discord.CommandInteractionType {
			commandInteraction := interaction.Data.(*discord.CommandInteraction)
			if result := commandHandler.Handle(ctx, commandInteraction, interaction); result != nil {
				response = result
			}
			return response
		}
		if interaction.Data.InteractionType() == discord.AutocompleteInteractionType {
			autoCompleteInteraction := interaction.Data.(*discord.AutocompleteInteraction)
			if result := autoCompleteHandler.Handle(ctx, autoCompleteInteraction); result != nil {
				response = result
			}
			return
		}
		if interaction.Data.InteractionType() == discord.ComponentInteractionType {
			componentInteraction := interaction.Data.(discord.ComponentInteraction)
			if result := componentHandler.Handle(ctx, componentInteraction, interaction); result != nil {
				response = result
			}
			return
		}
		if interaction.Data.InteractionType() == discord.ModalInteractionType {
			componentInteraction := interaction.Data.(*discord.ModalInteraction)
			if result := modalHandler.Handle(ctx, *componentInteraction, interaction); result != nil {
				response = result
			}
			return
		}
		return nil
	}()
}
