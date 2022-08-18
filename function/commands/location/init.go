package location

import (
	"function/base"
)

type LocationCommand struct{}
type LocationAutoCompleter struct{}

var _ base.Command = (*LocationCommand)(nil)
var _ base.AutoCompleter = (*LocationAutoCompleter)(nil)

func (c *LocationCommand) Name() string {
	return "location"
}

func (*LocationAutoCompleter) Name() string {
	return "location"
}
