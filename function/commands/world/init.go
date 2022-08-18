package world

import (
	"function/base"
)

type WorldCommand struct{}
type WorldComponent struct{}

var _ base.Command = (*WorldCommand)(nil)
var _ base.Component = (*WorldComponent)(nil)

func (c *WorldCommand) Name() string {
	return "world"
}

func (c *WorldComponent) Name() string {
	return "world"
}
