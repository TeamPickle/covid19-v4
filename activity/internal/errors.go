package internal

import (
	"activity/config"
	"errors"
	"runtime/debug"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
)

func SendInternalErrorMessage(recoverError any) {
	var err error
	if e, ok := recoverError.(error); ok {
		err = e
	} else if e, ok := recoverError.(string); ok {
		err = errors.New(e)
	} else {
		err = errors.New("unknown error")
	}
	client := webhook.New(config.LogWebhookID, config.LogWebhookToken)
	client.Execute(webhook.ExecuteData{
		Content: "activity 서버에서 오류가 발생했습니다.",
		Embeds: []discord.Embed{
			{Title: "Stack", Description: string(debug.Stack())},
			{Title: "Error", Description: err.Error()},
		},
	})
}
