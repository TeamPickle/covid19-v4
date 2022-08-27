package cron

import (
	"activity/external/coronaboard"
	"activity/internal"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
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
			if err := shard.(*state.State).Gateway().Send(context.Background(), &gateway.UpdatePresenceCommand{
				Activities: []discord.Activity{{Name: message}},
			}); err != nil {
				panic(err)
			}
		})
	}

	var getConfirmed func() string

	{
		lastFetched := time.Time{}
		currentConfirmed := ""
		getConfirmed = func() string {
			if time.Since(lastFetched) > time.Hour {
				lastFetched = time.Now()
				boardData, _ := coronaboard.ParseBoard(context.Background())
				currentConfirmed = humanize.Comma(boardData.ChartForGlobal.Kr.ConfirmedAcc[len(boardData.ChartForGlobal.Kr.ConfirmedAcc)-1])
			}
			return currentConfirmed
		}
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
		guilds := int64(0)
		m.ForEach(func(shard shard.Shard) {
			ss := shard.(*state.State)
			guildSlice, _ := ss.Guilds()
			guilds += int64(len(guildSlice))
		})
		changeActivity(fmt.Sprintf("%s개 서버에서 정보 확인", humanize.Comma(guilds)))
	})
	c.AddFunc("10,25,40,55 * * * * *", func() {
		changeActivity(fmt.Sprintf("현재 누적 확진자수 : %s명", getConfirmed()))
	})
	c.Start()
}
