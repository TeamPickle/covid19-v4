package world

import (
	"function/base"
	"function/config"

	"github.com/diamondburned/arikawa/v3/api"
)

type WorldCommand struct{}

var _ base.Command = (*WorldCommand)(nil)

func (c *WorldCommand) Name() string {
	return "world"
}

var client *api.Client

func init() {
	client = api.NewClient("Bot " + config.Token)
}
