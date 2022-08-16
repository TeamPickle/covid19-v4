//go:build dev

package main

import (
	"context"
	"encoding/json"
	"function/config"
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

	if interaction.Data.InteractionType() == discord.CommandInteractionType {
		commandInteraction := interaction.Data.(*discord.CommandInteraction)
		if result := commandHandler.Handle(ctx, commandInteraction); result != nil {
			return result
		}
	}

	return nil
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
