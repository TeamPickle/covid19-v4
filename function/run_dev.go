//go:build dev

package main

import (
	"context"
	"encoding/json"
	"function/config"
	"function/utils"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) *api.InteractionResponse {
	body := []byte(request.Body)
	var interaction discord.InteractionEvent
	interaction.UnmarshalJSON(body)

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
		return nil
	}()
}

func run() {
	s := session.New("Bot " + config.Token)
	s.AddHandler(func(e *gateway.InteractionCreateEvent) {
		jsonBytes, _ := json.Marshal(e)
		if resp := handler(context.Background(), events.LambdaFunctionURLRequest{Body: string(jsonBytes)}); resp != nil {
			s.RespondInteraction(e.ID, e.Token, *resp)
		}
	})
	s.AddIntents(gateway.IntentGuilds)

	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	log.Println("started")
	defer s.Close()
	select {}
}
