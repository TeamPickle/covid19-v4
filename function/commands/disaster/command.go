package disaster

import (
	"context"
	"encoding/json"
	"fmt"
	"function/utils"
	"io"
	"net/http"
	"regexp"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

var regions = []string{
	"전국", "강원", "경기", "경남", "경북", "광주", "대구", "대전", "부산", "서울", "울산", "인천", "전남", "전북", "제주", "충남", "충북", "세종",
}

func (*DisasterCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	region := data.Options.Find("province").String()
	regionIndex := -1
	for i, r := range regions {
		if r == region {
			regionIndex = i
			break
		}
	}

	client := http.Client{}
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://m.search.naver.com/p/csearch/content/nqapirender.nhn?where=m&pkid=258&key=disasterAlert&u1="+fmt.Sprintf("%02d", regionIndex), nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var bodyJson struct {
		Current struct {
			Html string `json:"html"`
		} `json:"current"`
	}
	json.Unmarshal(body, &bodyJson)

	searchLocalRegex := regexp.MustCompile(`<em class="area_name">(.+?)</em>`)
	searchConRegex := regexp.MustCompile(`<span class="dsc _text">(.+?)</span>`)
	searchDateTimeRegex := regexp.MustCompile(`<time datetime="">(.+?)</time>`)

	local := searchLocalRegex.FindAllStringSubmatch(bodyJson.Current.Html, -1)
	con := searchConRegex.FindAllStringSubmatch(bodyJson.Current.Html, -1)
	datetime := searchDateTimeRegex.FindAllStringSubmatch(bodyJson.Current.Html, -1)
	if len(local) == 0 && len(con) == 0 {
		return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
			Content: option.NewNullableString("재난문자를 불러올 수 없습니다."),
		})
	}

	embed := discord.Embed{
		Title:       "📌 재난문자",
		Description: fmt.Sprintf("%s 지역의 최근 5개 재난문자 목록입니다.", region),
		Color:       0xdd2255,
	}

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("%s(%s)", local[i][1], datetime[i][1])
		if len(name) > 100 {
			name = name[:99] + "…"
		}
		embed.Fields = append(embed.Fields, discord.EmbedField{
			Name:  name,
			Value: con[i][1],
		})
	}

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds: &[]discord.Embed{embed},
	})
}
