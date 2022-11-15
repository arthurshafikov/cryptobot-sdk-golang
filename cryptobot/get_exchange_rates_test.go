package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestGetExchangeRates(t *testing.T) {
	c := getTestClient(t)
	expectedExchangeRates := ExchangeRates{
		{
			IsValid: true,
			Rate:    "16836.58000000",
			Source:  "BTC",
			Target:  "USD",
		},
	}
	expectedResponse := getExchangeRatesResponse{
		response: response{
			Ok: true,
		},
		Result: expectedExchangeRates,
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getExchangeRates").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	currencies, err := c.GetExchangeRates()
	require.NoError(t, err)
	require.Equal(t, expectedExchangeRates, currencies)
}

func TestGetExchangeRatesReturnsError(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := getExchangeRatesResponse{
		response: response{
			Ok: false,
			Error: responseError{
				Code: 422,
				Name: "Unprocessable Entity",
			},
		},
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getExchangeRates").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	currencies, err := c.GetExchangeRates()
	require.Nil(t, currencies)
	require.ErrorContains(t, err, "getExchangeRates request error: code - 422, name - Unprocessable Entity")
}
