package status

import (
	"context"
	"fmt"
	"function/external/coronaboard"
	"function/utils"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/dustin/go-humanize"
)

func handleDomesticRegion(ctx context.Context, regionName string, boardData *coronaboard.CoronaBoardData) *api.InteractionResponse {
	var status *coronaboard.DomesticNowStatus
	for _, v := range boardData.StatDomesticNow {
		if v.Region == regionName {
			status = &v
			break
		}
	}
	if status == nil {
		panic("Can't find domestic region")
	}
	embed := discord.Embed{
		Title: fmt.Sprintf("시/도 확진자 수 조회 - %s", regionName),
		Description: fmt.Sprintf(""+
			"<:nujeok:687907310923677943> **확진자** : %s명(%s)\n"+
			"<:chiryojung:711728328985411616> **치료중** : %s명\n"+
			"<:wanchi:687907312052076594> 완치 : %s명\n"+
			"<:samang:687907312123510817> 사망 : %s명",
			humanize.Comma(status.Active),
			increase(status.Active-status.ActivePrev),
			humanize.Comma(status.Confirmed),
			humanize.Comma(status.Released),
			humanize.Comma(status.Death),
		),
	}

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds: &[]discord.Embed{embed},
	})
}

func handleForeignRegion(ctx context.Context, regionCode, regionName string, boardData *coronaboard.CoronaBoardData) *api.InteractionResponse {
	var status *coronaboard.GlobalNowStatus
	for _, v := range boardData.StatGlobalNow {
		if v.Cc == regionCode {
			status = &v
			break
		}
	}
	if status == nil {
		panic("Can't find foreign region")
	}

	embed := discord.Embed{
		Title: fmt.Sprintf("%s 국가별 현황 - %s", status.Flag, regionName),
		Description: fmt.Sprintf(""+
			"<:nujeok:687907310923677943> **확진자** : %s명(%s)\n"+
			"<:wanchi:687907312052076594> **완치자** : %s명(%s) - %d%%\n"+
			"<:samang:687907312123510817> **사망자** : %s명(%s) - %d%%",
			humanize.Comma(status.Confirmed),
			increase(status.Confirmed-status.ConfirmedPrev),
			humanize.Comma(status.Released),
			increase(status.Released-status.ReleasedPrev),
			(status.Released*100)/status.Confirmed,
			humanize.Comma(status.Death),
			increase(status.Death-status.DeathPrev),
			(status.Death*100)/status.Confirmed,
		),
		Color: 0x00bfff,
	}
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds: &[]discord.Embed{embed},
	})
}

func handleRegion(ctx context.Context, regionName string) *api.InteractionResponse {
	boardData, err := coronaboard.ParseBoard(ctx)
	if err != nil {
		panic("Can't parse board data")
	}
	regionName = aliasRegionName(regionName)
	regionCode := findRegionName(boardData.I18NAll.Ko, regionName)

	if isKoreanRegion(regionCode) {
		return handleDomesticRegion(ctx, regionName, boardData)
	}

	if regionCode != "" {
		return handleForeignRegion(ctx, regionCode, regionName, boardData)
	}

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString("" +
			"지원하지 않는 지역입니다." +
			"다음 중 하나를 입력해 주세요: `" + strings.Join(findKoreanRegions(boardData.I18NAll.Ko), "`, `") + "` 또는 `국가 이름`"),
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

func (c *StatusCommand) Handle(ctx context.Context, interaction *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	if len(interaction.Options) == 0 {
		return handleDomestic(ctx)
	}
	return handleRegion(ctx, interaction.Options[0].String())
}
