package world

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/dustin/go-humanize"
)

type mainData struct {
	active         int64
	confirmed      int64
	confirmedDelta int64
	released       int64
	releasedDelta  int64
	death          int64
	deathDelta     int64
	countries      int64
}

func makeMainEmbed(data *mainData) discord.Embed {
	return discord.Embed{
		Title: "🗺️ 세계 코로나 현황",
		Description: fmt.Sprintf(""+
			"<:chiryojung:711728328985411616> 치료중 : %s\n"+
			"<:nujeok:687907310923677943> 확진자 : %s(%s)\n"+
			"<:wanchi:687907312052076594> 완치 : %s(%s)\n"+
			"<:samang:687907312123510817> 사망 : %s(%s)\n"+
			"🚩 발생국 : %s",
			humanize.Comma(data.active),
			humanize.Comma(data.confirmed), humanize.Comma(data.confirmedDelta),
			humanize.Comma(data.released), humanize.Comma(data.releasedDelta),
			humanize.Comma(data.death), humanize.Comma(data.deathDelta),
			humanize.Comma(data.countries),
		),
		Color: 0x00cccc,
	}
}

func makeDetailedEmbed() discord.Embed {
	embed := discord.Embed{
		Title: "🗺️ 세계 코로나 현황",
		Color: 0x00cccc,
	}
	for _, status := range []string{"confirmed", "released", "death"} {
		embed.Description += fmt.Sprintf(""+
			"%s **%s** :\n"+
			"<:nujeok:687907310923677943> %s\n"+
			"<:wanchi:687907312052076594> %s\n"+
			"<:samang:687907312123510817> %s\n",
			status, "", "", "", "")
	}
	embed.Description += fmt.Sprintf("(1/%d)", 1)
	return embed
}
