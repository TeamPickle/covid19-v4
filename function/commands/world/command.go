package world

import (
	"context"
	"function/external/coronaboard"
	"function/utils"
	"math"
	"register/config"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (*WorldCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	go func() {
		coronaboardData, err := coronaboard.ParseBoard(ctx)
		if err != nil {
			resp := utils.MakeErrorMessage(rawRequest, "Can't get coronaboard data")
			client.EditInteractionResponse(config.AppID, rawRequest.Token, api.EditInteractionResponseData{Embeds: resp.Data.Embeds})
			return
		}
		client.EditInteractionResponse(config.AppID, rawRequest.Token, api.EditInteractionResponseData{
			Embeds: &[]discord.Embed{makeMainEmbed(&mainData{
				active:         coronaboardData.ChartForGlobal.Global.Active[len(coronaboardData.ChartForGlobal.Global.Active)-1],
				confirmed:      coronaboardData.ChartForGlobal.Global.ConfirmedAcc[len(coronaboardData.ChartForGlobal.Global.ConfirmedAcc)-1],
				confirmedDelta: coronaboardData.ChartForGlobal.Global.Confirmed[len(coronaboardData.ChartForGlobal.Global.Confirmed)-1],
				released:       coronaboardData.ChartForGlobal.Global.ReleasedAcc[len(coronaboardData.ChartForGlobal.Global.ReleasedAcc)-1],
				releasedDelta:  coronaboardData.ChartForGlobal.Global.Released[len(coronaboardData.ChartForGlobal.Global.Released)-1],
				death:          coronaboardData.ChartForGlobal.Global.DeathAcc[len(coronaboardData.ChartForGlobal.Global.DeathAcc)-1],
				deathDelta:     coronaboardData.ChartForGlobal.Global.Death[len(coronaboardData.ChartForGlobal.Global.Death)-1],
				countries:      int64(len(coronaboardData.StatGlobalNow)),
				totalPages:     int(math.Ceil(float64(len(coronaboardData.StatGlobalNow))/10)) + 1,
			})},
		})
	}()

	return &api.InteractionResponse{
		Type: api.DeferredMessageInteractionWithSource,
	}
}
