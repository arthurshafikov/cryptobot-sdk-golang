package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestGetCurrencies(t *testing.T) {
	c := getTestClient(t)
	expectedCurrencies := Currencies{
		{
			Code:         "BTC",
			Decimals:     8,
			IsBlockchain: true,
			Name:         "Bitcoin",
			Url:          "https://bitcoin.org/",
		},
		{
			Code:     "USD",
			Decimals: 8,
			IsFiat:   true,
			Name:     "United States dollar",
		},
	}
	expectedResponse := getCurrenciesResponse{
		response: response{
			Ok: true,
		},
		Result: expectedCurrencies,
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getCurrencies").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	currencies, err := c.GetCurrencies()
	require.NoError(t, err)
	require.Equal(t, expectedCurrencies, currencies)
}

func TestGetCurrenciesReturnsError(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := getCurrenciesResponse{
		response: response{
			Ok: false,
			Error: responseError{
				Code: 500,
				Name: "Server Error",
			},
		},
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getCurrencies").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	currencies, err := c.GetCurrencies()
	require.Nil(t, currencies)
	require.ErrorContains(t, err, "getCurrencies request error: code - 500, name - Server Error")
}
