package status

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (*StatusAutoCompleter) Handle(ctx context.Context, data *discord.AutocompleteInteraction) *api.InteractionResponse {
	boardData, err := parseBoard(ctx)
	if err != nil {
		panic("Can't parse board data")
	}
	partialRegion := strings.TrimSpace(data.Options.Find("region").String())
	choices := api.AutocompleteStringChoices{}
	for key, region := range boardData.I18NAll.Ko {
		if utf8.RuneCountInString(key) == 2 && strings.HasPrefix(region, partialRegion) {
			choices = append(choices, discord.StringChoice{Name: region, Value: region})
		}
		if len(choices) >= 25 {
			break
		}
	}

	return &api.InteractionResponse{
		Type: api.AutocompleteResult,
		Data: &api.InteractionResponseData{Choices: choices},
	}
}

func (*StatusAutoCompleter) Name() string {
	return "status"
}
