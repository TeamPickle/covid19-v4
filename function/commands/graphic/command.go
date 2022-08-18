package graphic

import (
	"context"
	"function/external/coronaboard"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/sendpart"
)

/*
const makeImage = async (data: ThenArg<ReturnType<typeof parseNcov>>) => {

  context.textBaseline = 'hanging';

*/

func (*GraphicCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	// resources.GraphicImages
	coronaboardData, err := coronaboard.ParseBoard(ctx)
	if err != nil {
		panic(err)
	}

	sum := int64(0)
	for _, v := range coronaboardData.StatDomesticNow {
		if v.Region == "검역" || v.Region == "합계" {
			continue
		}
		sum += v.Confirmed - v.ConfirmedPrev
	}

	dc := makeNewCanvas()
	fillBackground(dc)
	drawBackgroundImage(dc, "bg.png", 0, 0)
	for _, v := range coronaboardData.StatDomesticNow {
		if v.Region == "검역" || v.Region == "합계" {
			continue
		}
		per := float64(v.Confirmed-v.ConfirmedPrev) / float64(sum) * 100
		drawRegion(dc, per, v.Region)
	}
	drawBackgroundImage(dc, "k.png", 213, 10)

	for _, v := range coronaboardData.StatDomesticNow {
		if v.Region == "검역" || v.Region == "합계" {
			continue
		}
		drawRegionText(dc, v.Confirmed-v.ConfirmedPrev, v.Confirmed, v.Region)
	}

	reader := makeImage(dc)

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Files: []sendpart.File{{Name: "graphic.png", Reader: reader}},
	})
}
