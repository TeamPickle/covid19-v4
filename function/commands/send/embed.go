package send

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
)

type sendType string

const (
	속보   sendType = "속보"
	뉴스   sendType = "뉴스"
	해외   sendType = "해외"
	확진   sendType = "확진"
	사망   sendType = "사망"
	전체공지 sendType = "전체공지"
)

func (t sendType) String() string {
	return fmt.Sprintf("%s %s", t.Emoji(), string(t))
}

func (t sendType) Name() string {
	return string(t)
}

func (t sendType) Color() discord.Color {
	if t == 속보 {
		return 0xff4848
	}
	if t == 뉴스 {
		return 0x6699ff
	}
	if t == 해외 {
		return 0x9966ff
	}
	if t == 확진 {
		return 0xff7c80
	}
	if t == 사망 {
		return 0x222222
	}
	if t == 전체공지 {
		return 0x555555
	}
	return 0xffffff
}

func (t sendType) Emoji() string {
	if t == 속보 {
		return "<:sokbo:687907311875915845>"
	}
	if t == 뉴스 {
		return "<:gisa:687907312102670346>"
	}
	if t == 해외 {
		return "<:waeguk:687907310982791183>"
	}
	if t == 확진 {
		return "<:nujeok:687907310923677943>"
	}
	if t == 사망 {
		return "<:samang:687907312123510817>"
	}
	if t == 전체공지 {
		return "📢"
	}
	return ""
}

func makeEmbedWithContent(selectedType sendType, content string) *discord.Embed {
	return &discord.Embed{
		Title:       selectedType.String(),
		Description: content,
		Color:       selectedType.Color(),
	}
}
