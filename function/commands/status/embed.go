package status

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/dustin/go-humanize"
)

func increase(number int64) string {
	if number > 0 {
		return fmt.Sprintf("▲%s", humanize.Comma(number))
	}
	if number < 0 {
		return fmt.Sprintf("▼%s", humanize.Comma(number))
	}
	return "-0"
}

func makeEmbedWithData(data *nCovData, imageUrl string) *discord.Embed {
	return &discord.Embed{
		Title: fmt.Sprintf("대한민국 코로나19 확진 정보 (%s 기준)", data.date.Format("2006년 01월 02일 15시")),
		Description: fmt.Sprintf(""+
			"**확진자 현황**\n"+
			"<:case:905316440452792320> **누적 확진**: %s `전일 대비 %s`\n"+
			"<:death:905316440440176690> **사망**: %s `전일 대비 %s`",
			humanize.Comma(data.confirmedAccumulated),
			increase(data.confirmedDelta),
			humanize.Comma(data.deathAccumulated),
			increase(data.deathDelta),
		) + "\n\n" +
			"⚠️ 현재 국내 코로나19 확진 현황을 불러오는 공공 API가 매일 업데이트되지 않아, " +
			"이틀치의 확진자 및 사망자 합산치가 그래프에 기록되고 있습니다. " +
			"자세한 수치는 [질병관리청 보도자료](https://www.kdca.go.kr/board/board.es?mid=a20501010000&bid=0015) 에서 확인해주시기 바랍니다.\n" +
			"이용자 분들의 양해 부탁드립니다.",
		Color: 0x00669a,
		Footer: &discord.EmbedFooter{
			Text: "지자체에서 자체 집계한 자료와는 차이가 있을 수 있습니다.",
		},
		Image: &discord.EmbedImage{URL: imageUrl},
	}
}
