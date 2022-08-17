package base

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

type commandHandlerFunc func(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse

type Command interface {
	Handle(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse
	Name() string
}

type CommandHandler interface {
	Register(commands ...Command)
	Handle(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse
}

type commandHandler struct {
	handlers map[string]commandHandlerFunc
}

func NewCommandHandler() CommandHandler {
	return &commandHandler{handlers: map[string]commandHandlerFunc{}}
}

func (h *commandHandler) Register(commands ...Command) {
	for _, c := range commands {
		h.handlers[c.Name()] = c.Handle
	}
}

func (h *commandHandler) Handle(ctx context.Context, data *discord.CommandInteraction) *api.InteractionResponse {
	for k, hf := range h.handlers {
		if data.Name == k {
			return hf(ctx, data)
		}
	}
	return nil
}
