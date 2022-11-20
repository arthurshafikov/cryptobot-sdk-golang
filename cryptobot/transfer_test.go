package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

var expectedTransfer = Transfer{
	ID:          123,
	UserID:      "123123",
	Status:      "paid",
	Asset:       USDT,
	Amount:      "100",
	Comment:     "someComment",
	CompletedAt: "2022-11-24T05:29:38.495Z",
}

func TestTransfer(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := transferResponse{
		response: response{
			Ok: true,
		},
		Result: expectedTransfer,
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/transfer").
		MatchHeader(apiTokenHeaderName, testToken).
		MatchParam("asset", USDT).
		MatchParam("amount", "100").
		MatchParam("user_id", "123123").
		MatchParam("spend_id", "someSpendID").
		MatchParam("comment", "someComment").
		MatchParam("disable_send_notification", "true").
		Reply(200).
		JSON(expectedResponse)

	transfer, err := c.Transfer(TransferRequest{
		Asset:                   USDT,
		Amount:                  "100",
		UserID:                  123123,
		SpendID:                 "someSpendID",
		Comment:                 "someComment",
		DisableSendNotification: true,
	})
	require.NoError(t, err)
	require.Equal(t, expectedTransfer, *transfer)
}

func TestTransferReturnsError(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := transferResponse{
		response: response{
			Ok: false,
			Error: responseError{
				Code: 403,
				Name: "Forbidden",
			},
		},
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/transfer").
		MatchHeader(apiTokenHeaderName, testToken).
		MatchParam("asset", USDT).
		MatchParam("amount", "100").
		MatchParam("user_id", "123123").
		MatchParam("spend_id", "someSpendID").
		Reply(200).
		JSON(expectedResponse)

	transfer, err := c.Transfer(TransferRequest{
		Asset:   USDT,
		Amount:  "100",
		UserID:  123123,
		SpendID: "someSpendID",
	})
	require.Nil(t, transfer)
	require.ErrorContains(t, err, "transfer request error: code - 403, name - Forbidden")
}
