package cron

import (
	"activity/internal"
	"context"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/robfig/cron/v3"
)

type internalLogger struct{}

func (l internalLogger) Info(msg string, keysAndValues ...interface{}) {
}

func (l internalLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	internal.SendInternalErrorMessage(err)
}

func Start(s *session.Session) {
	c := cron.New(
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
		cron.WithChain(cron.Recover(internalLogger{})),
		cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))),
		cron.WithSeconds())
	c.AddFunc("0,15,30,45 * * * * *", func() {
		s.Gateway().Send(context.Background(), &gateway.UpdatePresenceCommand{
			Activities: []discord.Activity{{Name: "테스트1"}},
		})
	})
	c.AddFunc("5,20,35,50 * * * * *", func() {
		s.Gateway().Send(context.Background(), &gateway.UpdatePresenceCommand{
			Activities: []discord.Activity{{Name: "테스트2"}},
		})
	})
	c.AddFunc("10,25,40,55 * * * * *", func() {
		s.Gateway().Send(context.Background(), &gateway.UpdatePresenceCommand{
			Activities: []discord.Activity{{Name: "테스트3"}},
		})
	})
	c.Start()
}
