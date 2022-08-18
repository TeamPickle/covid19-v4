package base

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

type commandHandlerFunc func(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse

type Command interface {
	Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse
	Name() string
}

type CommandHandler interface {
	Register(commands ...Command)
	Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse
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

func (h *commandHandler) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	handler := h.handlers[data.Name]
	if handler == nil {
		return nil
	}
	return convertCustomID(data.Name, handler(ctx, data, rawRequest), rawRequest)
}
