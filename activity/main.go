package main

import (
	"activity/config"
	"activity/cron"
	"context"
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
)

func main() {
	newShard := state.NewShardFunc(func(m *shard.Manager, s *state.State) {
		s.AddIntents(gateway.IntentGuilds)
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
