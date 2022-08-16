package status

import "function/base"

type StatusCommand struct{}

var _ base.Command = (*StatusCommand)(nil)

func (c *StatusCommand) Name() string {
	return "status"
}
