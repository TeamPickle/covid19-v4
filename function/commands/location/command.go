package location

import (
	"context"
	"fmt"
	"function/database"
	"function/utils"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (*LocationCommand) Handle(ctx context.Context, data *discord.CommandInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	province := data.Options.Find("province")
	city := data.Options.Find("city")
	result := database.Location.FindOne(ctx, bson.M{"_id": rawRequest.SenderID().String()})
	userLocation := database.LocationUpdateProps{}
	if result.Err() == nil {
		result.Decode(&userLocation)
	}
	userLocation.Location = fmt.Sprintf("%s %s", province.String(), city.String())
	_, err := database.Location.ReplaceOne(ctx, bson.M{"_id": rawRequest.SenderID().String()}, &userLocation, options.Replace().SetUpsert(true))
	if err != nil {
		panic(err)
	}

	return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
		Content: option.NewNullableString(fmt.Sprintf("%s %s(으)로 지역을 설정했습니다.", province.String(), city.String())),
		Flags:   api.EphemeralResponse,
	})
}
