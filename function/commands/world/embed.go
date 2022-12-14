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
		Title: "πΊοΈ μΈκ³ μ½λ‘λ νν©",
		Description: fmt.Sprintf(""+
			"<:medicine:905916876796878859> μΉλ£μ€ : %s\n"+
			"<:case:905916877040136272> νμ§μ : %s(%s)\n"+
			"<:cure:905916877098876928> μμΉ : %s(%s)\n"+
			"<:death:905916876729778247> μ¬λ§ : %s(%s)\n"+
			"π© λ°μκ΅­ : %s",
			humanize.Comma(data.active),
			humanize.Comma(data.confirmed), humanize.Comma(data.confirmedDelta),
			humanize.Comma(data.released), humanize.Comma(data.releasedDelta),
			humanize.Comma(data.death), humanize.Comma(data.deathDelta),
			humanize.Comma(data.countries),
		),
		Color: 0x00cccc,
	}
}

type detailData struct {
	flag        string
	countryName string
	confirmed   int64
	released    int64
	death       int64
}

func makeDetailedEmbed(data []*detailData) discord.Embed {
	embed := discord.Embed{
		Title: "πΊοΈ μΈκ³ μ½λ‘λ νν©",
		Color: 0x00cccc,
	}
	for _, status := range data {
		embed.Description += fmt.Sprintf(""+
			"%s **%s** : "+
			"<:case:905916877040136272> %s / "+
			"<:cure:905916877098876928> %s /"+
			"<:death:905916876729778247> %s\n",
			status.flag, status.countryName,
			humanize.Comma(status.confirmed),
			humanize.Comma(status.released),
			humanize.Comma(status.death),
		)
	}
	return embed
}
