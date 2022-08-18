package channel

import (
	"function/base"
)

type ChannelCommand struct{}

var _ base.Command = (*ChannelCommand)(nil)

func (c *ChannelCommand) Name() string {
	return "channel"
}
