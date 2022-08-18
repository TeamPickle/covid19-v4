package world

import (
	"context"
	"fmt"
	"strconv"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func makeComponents(currentPage, totalPages int) *discord.ContainerComponents {
	return discord.ComponentsPtr(
		&discord.ActionRowComponent{
			&discord.ButtonComponent{
				Label:    "이전",
				Style:    discord.SecondaryButtonStyle(),
				CustomID: discord.ComponentID(fmt.Sprintf("%d", (currentPage+totalPages-1)%totalPages)),
			},
			&discord.ButtonComponent{
				Label:    fmt.Sprintf("페이지 (%d/%d)", currentPage+1, totalPages),
				Style:    discord.PrimaryButtonStyle(),
				CustomID: "page",
				Disabled: true,
			},
			&discord.ButtonComponent{
				Label:    "다음",
				Style:    discord.SecondaryButtonStyle(),
				CustomID: discord.ComponentID(fmt.Sprintf("%d", (currentPage+1)%totalPages)),
			},
		},
	)
}

// Handle implements base.Component
func (*WorldComponent) Handle(ctx context.Context, data discord.ComponentInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	nextPage, err := strconv.Atoi(string(data.ID()))
	if err != nil {
		panic(err)
	}
	if nextPage == 0 {
		embedData, totalPages := getMainData(ctx)

		return &api.InteractionResponse{
			Type: api.UpdateMessage,
			Data: &api.InteractionResponseData{
				Embeds:     &[]discord.Embed{makeMainEmbed(&embedData)},
				Components: makeComponents(0, totalPages),
			},
		}
	}

	detailDataList, totalPages := getListData(ctx, nextPage-1)
	if nextPage < 0 || nextPage >= totalPages {
		panic(fmt.Errorf("invalid page: %d", nextPage))
	}
	return &api.InteractionResponse{
		Type: api.UpdateMessage,
		Data: &api.InteractionResponseData{
			Embeds:     &[]discord.Embed{makeDetailedEmbed(detailDataList)},
			Components: makeComponents(nextPage, totalPages),
		},
	}
}
