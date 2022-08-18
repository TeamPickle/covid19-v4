package channel

import (
	"context"
	"fmt"
	"function/models"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (*ChannelCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	channel := data.Options.Find("channel")
	guildIDString := rawRequest.GuildID.String()
	settingUpdate := models.SettingUpdateProps{}
	models.Setting.FindOne(ctx, bson.M{"_id": guildIDString}).Decode(&settingUpdate)
	settingUpdate.Channel = channel.String()
	_, err := models.Setting.ReplaceOne(ctx, bson.M{"_id": guildIDString}, &settingUpdate, options.Replace().SetUpsert(true))
	if err != nil {
		panic(err)
	}
	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString(fmt.Sprintf("<#%s>(이)가 전체공지를 띄울 채널로 설정되었습니다.", channel.String())),
		Flags:   api.EphemeralResponse,
	})
}
