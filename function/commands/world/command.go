package world

import (
	"context"
	"fmt"
	"function/external/coronaboard"
	"function/utils"
	"math"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (*WorldCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	coronaboardData, err := coronaboard.ParseBoard(ctx)
	if err != nil {
		panic(err)
	}
	embedData := mainData{
		active:         coronaboardData.ChartForGlobal.Global.Active[len(coronaboardData.ChartForGlobal.Global.Active)-1],
		confirmed:      coronaboardData.ChartForGlobal.Global.ConfirmedAcc[len(coronaboardData.ChartForGlobal.Global.ConfirmedAcc)-1],
		confirmedDelta: coronaboardData.ChartForGlobal.Global.Confirmed[len(coronaboardData.ChartForGlobal.Global.Confirmed)-1],
		released:       coronaboardData.ChartForGlobal.Global.ReleasedAcc[len(coronaboardData.ChartForGlobal.Global.ReleasedAcc)-1],
		releasedDelta:  coronaboardData.ChartForGlobal.Global.Released[len(coronaboardData.ChartForGlobal.Global.Released)-1],
		death:          coronaboardData.ChartForGlobal.Global.DeathAcc[len(coronaboardData.ChartForGlobal.Global.DeathAcc)-1],
		deathDelta:     coronaboardData.ChartForGlobal.Global.Death[len(coronaboardData.ChartForGlobal.Global.Death)-1],
		countries:      int64(len(coronaboardData.StatGlobalNow)),
		totalPages:     int(math.Ceil(float64(len(coronaboardData.StatGlobalNow))/10)) + 1,
	}
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds: &[]discord.Embed{makeMainEmbed(&embedData)},
		Components: discord.ComponentsPtr(
			&discord.ActionRowComponent{
				&discord.ButtonComponent{
					Style:    discord.PrimaryButtonStyle(),
					CustomID: discord.ComponentID(fmt.Sprintf("%s-%d", rawRequest.SenderID(), embedData.totalPages-1)),
					Emoji:    &discord.ComponentEmoji{Name: "◀️"},
				},
				&discord.ButtonComponent{
					Style:    discord.PrimaryButtonStyle(),
					CustomID: discord.ComponentID(fmt.Sprintf("%s-%d", rawRequest.SenderID(), 1)),
					Emoji:    &discord.ComponentEmoji{Name: "▶️"},
				},
			},
		),
	})
}
