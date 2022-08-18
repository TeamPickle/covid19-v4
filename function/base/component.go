package base

import (
	"context"
	"fmt"
	"function/utils"
	"reflect"
	"strings"
	"unsafe"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

type componentHandlerFunc func(ctx context.Context, data discord.ComponentInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse

type Component interface {
	Handle(ctx context.Context, data discord.ComponentInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse
	Name() string
}

type ComponentHandler interface {
	Register(commands ...Component)
	Handle(ctx context.Context, data discord.ComponentInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse
}

type componentHandler struct {
	handlers map[string]componentHandlerFunc
}

func NewComponentHandler() ComponentHandler {
	return &componentHandler{handlers: map[string]componentHandlerFunc{}}
}

func (h *componentHandler) Register(commands ...Component) {
	for _, c := range commands {
		h.handlers[c.Name()] = c.Handle
	}
}

func (h *componentHandler) Handle(ctx context.Context, data discord.ComponentInteraction, rawRequest discord.InteractionEvent) *api.InteractionResponse {
	tokens := strings.Split(string(data.ID()), ":")
	if len(tokens) != 3 {
		return nil
	}
	name := tokens[0]
	userIDSnowflake, _ := discord.ParseSnowflake(tokens[1])
	userID := discord.UserID(userIDSnowflake)
	action := tokens[2]

	if userID != rawRequest.SenderID() {
		return utils.MessageInteractionResponseWithSource(&api.InteractionResponseData{
			Content: option.NewNullableString(fmt.Sprintf("/%s 명령어를 직접 입력해서 사용해주세요", rawRequest.Message.Interaction.Name)),
			Flags:   api.EphemeralResponse,
		})
	}

	handler := h.handlers[name]
	if handler == nil {
		return nil
	}

	dataValue := reflect.ValueOf(data)
	dataIndirect := reflect.Indirect(dataValue)
	idField := dataIndirect.FieldByName("id")
	pointer := unsafe.Pointer(idField.Addr().Pointer())
	idPointer := (*string)(pointer)
	*idPointer = action

	return convertCustomID(name, handler(ctx, data, rawRequest), rawRequest)
}
