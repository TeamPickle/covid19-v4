package main

import (
	"function/base"
	"function/commands/channel"
	"function/commands/disaster"
	"function/commands/graphic"
	"function/commands/location"
	"function/commands/send"
	"function/commands/status"
	"function/commands/world"
)

var (
	commandHandler      = base.NewCommandHandler()
	autoCompleteHandler = base.NewAutoCompleteHandler()
	componentHandler    = base.NewComponentHandler()
	modalHandler        = base.NewModalHandler()
)

func main() {
	commandHandler.Register(
		&status.StatusCommand{},
		&disaster.DisasterCommand{},
		&world.WorldCommand{},
		&location.LocationCommand{},
		&graphic.GraphicCommand{},
		&channel.ChannelCommand{},
		&send.SendCommand{},
	)
	autoCompleteHandler.Register(
		&status.StatusAutoCompleter{},
		&location.LocationAutoCompleter{},
	)
	componentHandler.Register(
		&world.WorldComponent{},
		&send.SendComponent{},
	)
	modalHandler.Register(
		&send.SendModal{},
	)

	run()
}
