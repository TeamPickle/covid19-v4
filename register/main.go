package main

import (
	"fmt"
	"log"
	"register/config"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

var (
	client *api.Client
)

var provinceOption = discord.StringOption{
	OptionName: "province",
	OptionNameLocalizations: discord.StringLocales{
		discord.Korean: "시도",
	},
	Description: "시/도를 입력하세요. 예시: 서울, 경기도",
	Choices: []discord.StringChoice{
		{Name: "강원"},
		{Name: "경기"},
		{Name: "경남"},
		{Name: "경상남도", Value: "경남"},
		{Name: "경북"},
		{Name: "경상북도", Value: "경북"},
		{Name: "광주"},
		{Name: "대구"},
		{Name: "대전"},
		{Name: "부산"},
		{Name: "서울"},
		{Name: "울산"},
		{Name: "인천"},
		{Name: "전남"},
		{Name: "전라남도", Value: "전남"},
		{Name: "전북"},
		{Name: "전라북도", Value: "전북"},
		{Name: "제주"},
		{Name: "충남"},
		{Name: "충청남도", Value: "충남"},
		{Name: "충북"},
		{Name: "충청북도", Value: "충북"},
		{Name: "세종"},
	},
	Required: true,
}

func init() {
	client = api.NewClient("Bot " + config.Token)
}

func addCommand(data api.CreateCommandData) {
	if config.TestGuildId == discord.NullGuildID {
		if _, err := client.CreateCommand(config.AppID, data); err != nil {
			panic(err)
		}
	} else {
		if command, err := client.CreateGuildCommand(config.AppID, config.TestGuildId, data); err != nil {
			panic(err)
		} else {
			fmt.Println(command)
		}
	}
}

func main() {
	log.Println("Registering commands...")
	addCommand(api.CreateCommandData{
		Name: "status",
		NameLocalizations: discord.StringLocales{
			discord.Korean: "현황",
		},
		Description: "현황을 불러옵니다.",
		Options: discord.CommandOptions{&discord.StringOption{
			OptionName: "region",
			OptionNameLocalizations: discord.StringLocales{
				discord.Korean: "지역",
			},
			Description:  "불러오고 싶은 시/도, 또는 국가를 입력하세요. 예시: 서울, 한국",
			Required:     false,
			Autocomplete: true,
		}},
	})
	addCommand(api.CreateCommandData{
		Name: "graphic",
		NameLocalizations: discord.StringLocales{
			discord.Korean: "국내현황",
		},
		Description: "전국 시/도의 현황을 불러와 사진으로 보여줍니다",
	})
	addCommand(api.CreateCommandData{
		Name: "world",
		NameLocalizations: discord.StringLocales{
			discord.Korean: "세계현황",
		},
		Description: "전 세계 국가의 현황을 불러옵니다.",
	})
	addCommand(api.CreateCommandData{
		Name: "location",
		NameLocalizations: discord.StringLocales{
			discord.Korean: "위치설정",
		},
		Description: "내 위치를 등록합니다. 등록 후에는 명령어에 지역을 입력하지 않아도 등록한 위치의 정보를 불러옵니다.",
		Options: discord.CommandOptions{
			&provinceOption,
			&discord.StringOption{
				OptionName: "city",
				OptionNameLocalizations: discord.StringLocales{
					discord.Korean: "시군구",
				},
				Description:  "시/군/구를 입력하세요. 예시: 광진구, 수원시",
				Autocomplete: true,
			},
		},
	})
	addCommand(api.CreateCommandData{
		Name: "channel",
		NameLocalizations: discord.StringLocales{
			discord.Korean: "채널설정",
		},
		Description: "뉴스 및 전체공지를 전송하는 채널을 설정합니다.",
		Options: discord.CommandOptions{
			discord.NewChannelOption("name", "채널명", true),
		},
		DefaultMemberPermissions: discord.NewPermissions(discord.PermissionAdministrator),
	})
	addCommand(api.CreateCommandData{
		Name:                     "disaster",
		NameLocalizations:        discord.StringLocales{discord.Korean: "재난문자"},
		Description:              "특정 지역의 재난문자를 불러옵니다.",
		Options:                  discord.CommandOptions{&provinceOption},
		DefaultMemberPermissions: discord.NewPermissions(discord.PermissionAdministrator),
	})
	log.Println("Done.")
}
