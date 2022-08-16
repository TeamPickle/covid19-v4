package base

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

type handlerFunc func(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse

type Command struct {
	Name    string
	Handler handlerFunc
}

type CommandHandler interface {
	Register(commands ...*Command)
	Handle(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse
}

type commandHandler struct {
	handlers map[string]handlerFunc
}

func NewCommandHandler() CommandHandler {
	return &commandHandler{}
}

func (h *commandHandler) Register(commands ...*Command) {

}

func (h *commandHandler) Handle(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse {
	for k, hf := range h.handlers {
		if data.Name == k {
			return hf(ctx, data)
		}
	}
	return nil
}
