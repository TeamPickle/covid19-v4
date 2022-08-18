package base

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func convertComponent(components *discord.ContainerComponents, prefix string) *discord.ContainerComponents {
	if components == nil {
		return nil
	}
	for _, component := range *components {
		if component.Type() == discord.ActionRowComponentType {
			actionRowComponent := component.(*discord.ActionRowComponent)
			for _, action := range *actionRowComponent {
				if action.Type() == discord.ButtonComponentType {
					buttonComponent := action.(*discord.ButtonComponent)
					buttonComponent.CustomID = discord.ComponentID(prefix + string(buttonComponent.CustomID))
				} else if action.Type() == discord.SelectComponentType {
					selectComponent := action.(*discord.SelectComponent)
					selectComponent.CustomID = discord.ComponentID(prefix + string(selectComponent.CustomID))
				}
			}
		}
	}
	return components
}

func convertCustomID(name string, result *api.InteractionResponse, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	userID := rawRequest.SenderID()
	customIDPrefix := fmt.Sprintf("%s:%s:", name, userID)
	if result == nil || result.Data == nil {
		return result
	}
	if result.Data.CustomID != nil {
		result.Data.CustomID.Val = name + ":" + result.Data.CustomID.Val
	}
	convertComponent(result.Data.Components, customIDPrefix)
	return result
}
