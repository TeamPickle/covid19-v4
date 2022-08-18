package graphic

import (
	"function/base"
)

type GraphicCommand struct{}

var _ base.Command = (*GraphicCommand)(nil)

func (c *GraphicCommand) Name() string {
	return "graphic"
}
