package disaster

import (
	"function/base"
)

type DisasterCommand struct{}

var _ base.Command = (*DisasterCommand)(nil)

func (c *DisasterCommand) Name() string {
	return "disaster"
}
