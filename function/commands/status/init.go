package status

import (
	"function/base"
)

type StatusCommand struct{}
type StatusAutoCompleter struct{}

var _ base.Command = (*StatusCommand)(nil)
var _ base.AutoCompleter = (*StatusAutoCompleter)(nil)

func (c *StatusCommand) Name() string {
	return "status"
}

func (*StatusAutoCompleter) Name() string {
	return "status"
}
