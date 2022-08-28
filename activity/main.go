package main

import (
	"activity/config"
	"activity/cron"
	"context"
	"log"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/diamondburned/arikawa/v3/utils/ws"
)

func main() {
	ws.WSDebug = func(v ...any) {
		log.Print("WS: ")
		log.Println(v)
	}
	newShard := state.NewShardFunc(func(m *shard.Manager, s *state.State) {
		s.AddIntents(gateway.IntentGuilds)
		s.AddHandler(func(e *gateway.InteractionCreateEvent) {
			if e.Data.InteractionType() != discord.CommandInteractionType {
				return
			}
			cmdInteraction := e.Data.(*discord.CommandInteraction)
			if cmdInteraction.Name != "info" {
				return
			}
			s.RespondInteraction(e.ID, e.Token, api.InteractionResponse{
				Type: api.MessageInteractionWithSource,
				Data: &api.InteractionResponseData{
					Content: option.NewNullableString("I'm alive!"),
				},
			})
		})
	})
	m, err := shard.NewManager("Bot "+config.Token, newShard)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	defer m.Close()

	go cron.Start(m)
	runWebServer(m)
}
