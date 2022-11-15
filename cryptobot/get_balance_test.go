package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {
	c := getTestClient(t)
	expectedBalance := Balance{
		{
			Available:    "12",
			CurrencyCode: "USDT",
		},
		{
			Available:    "0.0026",
			CurrencyCode: "BTC",
		},
	}
	expectedResponse := getBalanceResponse{
		response: response{
			Ok: true,
		},
		Result: expectedBalance,
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getBalance").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	balance, err := c.GetBalance()
	require.NoError(t, err)
	require.Equal(t, expectedBalance, balance)
}

func TestGetBalanceReturnsError(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := getBalanceResponse{
		response: response{
			Ok: false,
			Error: responseError{
				Code: 401,
				Name: "Unauthorized",
			},
		},
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getBalance").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	balance, err := c.GetBalance()
	require.Nil(t, balance)
	require.ErrorContains(t, err, "getBalance request error: code - 401, name - Unauthorized")
}
