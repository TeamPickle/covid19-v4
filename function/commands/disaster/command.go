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

/*
const location = await getLocation(msg.author.id, args[1]);
    if (!location) {
      return msg.respond(stripIndents`
        명령어 사용법 : \`${getGuildPrefix(msg.guild)}재난문자 [지역]\`
        지역 목록 : \`${disasterData.disasterRegion.join(' ')}\`
      `);
    }
    const u = (
      Object.keys(disasterData.disasterAlias).includes(location)
        ? disasterData.disasterAlias[
            location as keyof typeof disasterData.disasterAlias
          ]
        : location
    ) as typeof disasterData.disasterRegion[number];

    const disasterIndex = disasterData.disasterRegion.indexOf(u);

    if (disasterIndex < 0) {
      return msg.respond(stripIndents`
        지원하지 않는 지역입니다. 다음 지역 중 하나로 다시 시도해주세요.
        \`${disasterData.disasterRegion.join(' ')}\`
      `);
    }

    const source: string = (
      await (
        await fetch(
          `https://m.search.naver.com/p/csearch/content/nqapirender.nhn?where=m&pkid=258&key=disasterAlert&u1=${
            disasterIndex ? disasterIndex.toString().padStart(2, '0') : ''
          }`,
        )
      ).json()
    ).current.html;

    const local = [...source.matchAll(/<em class="area_name">(.+?)<\/em>/g)];
    const con = [...source.matchAll(/<span class="dsc _text">(.+?)<\/span>/g)];
    const distime = [...source.matchAll(/<time datetime="">(.+?)<\/time>/g)];

    if (local.length === 0 && con.length === 0) {
      return msg.respond('재난문자를 불러올 수 없습니다.');
    }

    const embed = new MessageEmbed();
    embed
      .setTitle('📌 재난문자')
      .setDescription(`${u} 지역의 최근 5개 재난문자 목록입니다.`)
      .setColor(0xdd2255)
      .addFields(
        [...Array(5)].map((_, i) => ({
          name: `${local[i][1]}(${distime[i][1]})`.slice(0, 256),
          value: con[i][1],
        })),
      );
    return msg.respond({ embeds: [embed] });
*/

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
		embed.Fields = append(embed.Fields, discord.EmbedField{
			Name:  fmt.Sprintf("%s(%s)", local[i][1], datetime[i][1]),
			Value: con[i][1],
		})
	}

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Embeds: &[]discord.Embed{embed},
	})
}
