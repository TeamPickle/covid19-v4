//go:build !dev

package main

import (
	"context"
	"function/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	if !validateSignature(request) {
		return events.LambdaFunctionURLResponse{StatusCode: 401, Body: "Signature verification failed"}, nil
	}

	body := []byte(request.Body)
	var interaction discord.InteractionEvent
	interaction.UnmarshalJSON(body)

	if interaction.Data.InteractionType() == discord.PingInteractionType {
		return responseOK(&api.InteractionResponse{Type: api.PongInteraction})
	}

	response := func() (response *api.InteractionResponse) {
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
			return
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
	return responseOK(response)
}

func run() {
	lambda.Start(handler)
}
