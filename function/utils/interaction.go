package utils

import "github.com/diamondburned/arikawa/v3/api"

func MessageInteractionResponseWithSource(data *api.InteractionResponseData) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: data,
	}
}
