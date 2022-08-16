package main

import (
	"function/base"
	"function/commands/status"
)

var commandHandler base.CommandHandler

func main() {
	commandHandler = base.NewCommandHandler()
	commandHandler.Register(
		&status.StatusCommand{},
	)
	run()
}
