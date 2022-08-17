package main

import (
	"function/base"
	"function/commands/status"
)

var (
	commandHandler      base.CommandHandler
	autoCompleteHandler base.AutoCompleteHandler
)

func main() {
	commandHandler = base.NewCommandHandler()
	commandHandler.Register(
		&status.StatusCommand{},
	)
	autoCompleteHandler = base.NewAutoCompleteHandler()
	autoCompleteHandler.Register(
		&status.StatusAutoCompleter{},
	)
	run()
}
