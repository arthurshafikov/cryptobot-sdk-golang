package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestGetMe(t *testing.T) {
	c := getTestClient(t)
	expectedAppInfo := AppInfo{
		AppID:                        123,
		Name:                         "someName",
		PaymentProcessingBotUsername: "CryptoBot",
	}
	expectedResponse := getMeResponse{
		response: response{
			Ok: true,
		},
		Result: expectedAppInfo,
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getMe").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	appInfo, err := c.GetMe()
	require.NoError(t, err)
	require.Equal(t, expectedAppInfo, *appInfo)
}
