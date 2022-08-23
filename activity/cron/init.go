package cron

import (
	"activity/internal"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/session/shard"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/dustin/go-humanize"
	"github.com/robfig/cron/v3"
)

type internalLogger struct{}

func (l internalLogger) Info(msg string, keysAndValues ...interface{}) {
}

func (l internalLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	internal.SendInternalErrorMessage(err)
}

func Start(m *shard.Manager) {
	changeActivity := func(message string) {
		m.ForEach(func(shard shard.Shard) {
			shard.(*session.Session).Gateway().Send(context.Background(), &gateway.UpdatePresenceCommand{
				Activities: []discord.Activity{{Name: message}},
			})
		})
	}
	c := cron.New(
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
		cron.WithChain(cron.Recover(internalLogger{})),
		cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))),
		cron.WithSeconds())
	c.AddFunc("0,15,30,45 * * * * *", func() {
		changeActivity("/도움 으로 명령어 확인")
	})
	c.AddFunc("5,20,35,50 * * * * *", func() {
		guilds := 0
		m.ForEach(func(shard shard.Shard) {
			ss := shard.(*state.State)
			guildSlice, _ := ss.Guilds()
			guilds += len(guildSlice)
		})
		changeActivity(fmt.Sprintf("%d개 서버에서 정보 확인", guilds))
	})
	c.AddFunc("10,25,40,55 * * * * *", func() {
		// boardData, _ := coronaboard.ParseBoard(context.Background())
		changeActivity(fmt.Sprintf(
			"현재 누적 확진자수 : %s명",
			humanize.Comma(
				12345,
				// boardData.ChartForGlobal.Kr.ConfirmedAcc[len(boardData.ChartForGlobal.Kr.ConfirmedAcc)-1],
			),
		))
	})
	c.Start()
}
