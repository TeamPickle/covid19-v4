package send

import (
	"function/base"
)

type SendCommand struct{}
type SendModal struct{}
type SendComponent struct{}

var _ base.Command = (*SendCommand)(nil)
var _ base.Modal = (*SendModal)(nil)
var _ base.Component = (*SendComponent)(nil)

func (c *SendCommand) Name() string {
	return "send"
}

func (c *SendModal) Name() string {
	return "send"
}

func (*SendComponent) Name() string {
	return "send"
}
