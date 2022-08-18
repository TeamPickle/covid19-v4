package main

import (
	"function/base"
	"function/commands/disaster"
	"function/commands/graphic"
	"function/commands/location"
	"function/commands/status"
	"function/commands/world"
	"function/database"
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
		&location.LocationCommand{},
		&graphic.GraphicCommand{},
	)
	autoCompleteHandler = base.NewAutoCompleteHandler()
	autoCompleteHandler.Register(
		&status.StatusAutoCompleter{},
		&location.LocationAutoCompleter{},
	)
	componentHandler = base.NewComponentHandler()
	componentHandler.Register(
		&world.WorldComponent{},
	)
	database.Connect()
	run()
}
