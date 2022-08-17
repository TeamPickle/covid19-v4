package status

import (
	"context"
	"function/utils"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func handleRegion(ctx context.Context, regionName string) *api.InteractionResponse {
	boardData, err := parseBoard()
	if err != nil {
		panic("Can't parse board data")
	}
	regionName = aliasRegionName(regionName)
	regionCode := findRegionName(boardData.I18NAll.Ko, regionName)

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString("지역 현황 " + regionName + regionCode + "  " + strings.Join(findKoreanRegions(boardData.I18NAll.Ko), ", ")),
	})
}

func handleDomestic(ctx context.Context) *api.InteractionResponse {
	ncovData, err := parseNCov(ctx)
	if err != nil {
		panic("Can't parse ncov data")
	}
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString("국내 현황"),
		Embeds: &[]discord.Embed{
			*makeEmbedWithData(*ncovData[0], generateChartURL(ncovData)),
		},
	})
}

func (c *StatusCommand) Handle(ctx context.Context, interaction *discord.CommandInteraction) *api.InteractionResponse {
	if len(interaction.Options) == 0 {
		return handleDomestic(ctx)
	}
	return handleRegion(ctx, interaction.Options[0].String())
}
