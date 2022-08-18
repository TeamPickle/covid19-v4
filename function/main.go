package main

import (
	"function/base"
	"function/commands/disaster"
	"function/commands/status"
	"function/commands/world"
)

var (
	commandHandler      base.CommandHandler
	autoCompleteHandler base.AutoCompleteHandler
	componentHandler    base.ComponentHandler
)

func main() {
	commandHandler = base.NewCommandHandler()
	commandHandler.Register(
		&status.StatusCommand{},
		&disaster.DisasterCommand{},
		&world.WorldCommand{},
	)
	autoCompleteHandler = base.NewAutoCompleteHandler()
	autoCompleteHandler.Register(
		&status.StatusAutoCompleter{},
	)
	componentHandler = base.NewComponentHandler()
	componentHandler.Register(
		&world.WorldComponent{},
	)
	run()
}
