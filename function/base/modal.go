package base

import (
	"context"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

type modalHandlerFunc func(ctx context.Context, data discord.ModalInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse

type Modal interface {
	Handle(ctx context.Context, data discord.ModalInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse
	Name() string
}

type ModalHandler interface {
	Register(commands ...Modal)
	Handle(ctx context.Context, data discord.ModalInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse
}

type modalHandler struct {
	handlers map[string]modalHandlerFunc
}

func NewModalHandler() ModalHandler {
	return &modalHandler{handlers: map[string]modalHandlerFunc{}}
}

func (h *modalHandler) Register(commands ...Modal) {
	for _, c := range commands {
		h.handlers[c.Name()] = c.Handle
	}
}

func (h *modalHandler) Handle(ctx context.Context, data discord.ModalInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	tokens := strings.Split(string(data.CustomID), ":")
	if len(tokens) != 2 {
		return nil
	}
	name := tokens[0]
	action := tokens[1]

	handler := h.handlers[name]
	if handler == nil {
		return nil
	}

	data.CustomID = discord.ComponentID(action)

	return convertCustomID(name, handler(ctx, data, rawRequest), rawRequest)
}
