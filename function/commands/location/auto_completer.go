package location

import (
	"context"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (*LocationAutoCompleter) Handle(ctx context.Context, data *discord.AutocompleteInteraction) *api.InteractionResponse {
	province := data.Options.Find("province")
	city := data.Options.Find("city")
	cityStr := strings.TrimSpace(city.String())
	cities := regions[province.String()]
	choices := api.AutocompleteStringChoices{}

	for _, city := range cities {
		if strings.HasPrefix(city, cityStr) {
			choices = append(choices, discord.StringChoice{
				Name:  city,
				Value: city,
			})
		}
	}

	if len(choices) > 25 {
		choices = choices[:25]
	}

	return &api.InteractionResponse{
		Type: api.AutocompleteResult,
		Data: &api.InteractionResponseData{Choices: choices},
	}
}
