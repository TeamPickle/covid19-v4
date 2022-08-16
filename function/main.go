package main

import (
	"context"
	"function/base"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

var commandHandler base.CommandHandler

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

	if interaction.Data.InteractionType() == discord.CommandInteractionType {
		commandInteraction := interaction.Data.(*discord.CommandInteraction)
		if result := commandHandler.Handle(ctx, commandInteraction); result != nil {
			return responseOK(result)
		}
	}

	return events.LambdaFunctionURLResponse{StatusCode: 404}, nil
}

func main() {
	commandHandler = base.NewCommandHandler()
	commandHandler.Register()
	lambda.Start(handler)
}
