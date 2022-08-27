package main

import (
	"context"
	"function/config"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

func run() {
	if config.PublicKey != "" {
		lambdaHandler := func(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
			if !validateSignature(request) {
				return events.LambdaFunctionURLResponse{StatusCode: 401, Body: "Signature verification failed"}, nil
			}

			body := []byte(request.Body)
			var interaction discord.InteractionEvent
			interaction.UnmarshalJSON(body)

			res := handler(ctx, interaction)
			return responseOK(res)
		}
		lambda.Start(lambdaHandler)
	} else {
		s := session.New("Bot " + config.Token)
		s.AddHandler(func(e *gateway.InteractionCreateEvent) {
			if resp := handler(context.Background(), e.InteractionEvent); resp != nil {
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
}
