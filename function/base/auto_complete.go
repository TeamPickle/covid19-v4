package base

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

type autocompleteHandlerFunc func(ctx context.Context, data *discord.AutocompleteInteraction) *api.InteractionResponse

type AutoCompleter interface {
	Handle(ctx context.Context, data *discord.AutocompleteInteraction) *api.InteractionResponse
	Name() string
}

type AutoCompleteHandler interface {
	Register(commands ...AutoCompleter)
	Handle(ctx context.Context, data *discord.AutocompleteInteraction) *api.InteractionResponse
}

type autoCompleteHandler struct {
	handlers map[string]autocompleteHandlerFunc
}

func NewAutoCompleteHandler() AutoCompleteHandler {
	return &autoCompleteHandler{handlers: map[string]autocompleteHandlerFunc{}}
}

func (h *autoCompleteHandler) Register(commands ...AutoCompleter) {
	for _, c := range commands {
		h.handlers[c.Name()] = c.Handle
	}
}

func (h *autoCompleteHandler) Handle(ctx context.Context, data *discord.AutocompleteInteraction) *api.InteractionResponse {
	for k, hf := range h.handlers {
		if data.Name == k {
			return hf(ctx, data)
		}
	}
	return nil
}
