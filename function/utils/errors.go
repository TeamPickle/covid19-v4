package utils

import (
	"errors"
	"fmt"
	"function/config"
	"math/rand"
	"runtime/debug"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/martinlindhe/base36"
)

func MakeErrorMessage(req discord.InteractionEvent, recoverError any) *api.InteractionResponse {
	var err error
	if e, ok := recoverError.(error); ok {
		err = e
	} else if e, ok := recoverError.(string); ok {
		err = errors.New(e)
	} else {
		err = errors.New("unknown error")
	}
	client := webhook.New(config.LogWebhookID, config.LogWebhookToken)
	codeBytes := make([]byte, 8)
	rand.Read(codeBytes)
	code := strings.ToLower(base36.EncodeBytes(codeBytes))
	commandName := ""

	if req.Data.InteractionType() == discord.CommandInteractionType {
		command := req.Data.(*discord.CommandInteraction)
		arguments := ""
		for _, option := range command.Options {
			arguments += fmt.Sprintf(" [%s: %s]", option.Name, option.Value.String())
		}
		commandName = fmt.Sprintf("%s%s", command.Name, arguments)
	}

	client.Execute(webhook.ExecuteData{
		Content: fmt.Sprintf(""+
			"오류가 발생했습니다.\n"+
			"Author: %s (%s)\n"+
			"Command: %s\n"+
			"Code: %s\n",
			req.Sender().Tag(),
			req.SenderID(),
			commandName,
			code,
		),
		Embeds: []discord.Embed{
			{Title: "Stack", Description: string(debug.Stack())},
			{Title: "Error", Description: err.Error()},
		},
	})

	return MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds: &[]discord.Embed{
			{
				Title: "오류",
				Description: fmt.Sprintf(""+
					"오류가 발생했습니다.\n"+
					"[포럼](https://forum.tpk.kr)에서 오류를 제보해주세요.\n"+
					"코드: %s", code,
				),
				Color: 0xff3333,
			},
		},
	})
}
