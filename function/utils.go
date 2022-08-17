package main

import (
	"encoding/hex"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/diamondburned/arikawa/v3/api"
	"golang.org/x/crypto/nacl/sign"
)

func validateSignature(request events.LambdaFunctionURLRequest) bool {
	signature := request.Headers["x-signature-ed25519"]
	timestamp := request.Headers["x-signature-timestamp"]
	strBody := request.Body

	signedMessage, _ := hex.DecodeString(signature)
	publicKeyByteSlice, _ := hex.DecodeString("<PUBLIC KEY>")

	signedMessage = append(signedMessage, []byte(timestamp+strBody)...)
	_, isVerified := sign.Open(nil, signedMessage, (*[32]byte)(publicKeyByteSlice))
	return isVerified
}

func responseOK(data *api.InteractionResponse) (events.LambdaFunctionURLResponse, error) {
	if data == nil {
		return events.LambdaFunctionURLResponse{StatusCode: 404}, nil
	}
	messageBytes, _ := json.Marshal(data)
	return events.LambdaFunctionURLResponse{StatusCode: 200, Body: string(messageBytes)}, nil
}
