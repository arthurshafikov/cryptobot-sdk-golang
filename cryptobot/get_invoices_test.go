package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestGetInvoices(t *testing.T) {
	c := getTestClient(t)
	secondInvoice := expectedInvoice
	secondInvoice.ID = 22
	expectedInvoices := []Invoice{
		expectedInvoice,
		secondInvoice,
	}
	expectedResponse := getInvoicesResponse{
		response: response{
			Ok: true,
		},
		Result: invoices{
			Items: expectedInvoices,
		},
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/getInvoices").
		MatchHeader(apiTokenHeaderName, testToken).
		MatchParam("asset", "USDT").
		MatchParam("invoice_ids", "1,2,3").
		MatchParam("status", "paid").
		MatchParam("offset", "2").
		MatchParam("count", "100").
		Reply(200).
		JSON(expectedResponse)

	invoices, err := c.GetInvoices(&GetInvoicesRequest{
		Asset:      "USDT",
		InvoiceIDs: "1,2,3",
		Status:     "paid",
		Offset:     2,
		Count:      100,
	})
	require.NoError(t, err)
	require.Equal(t, expectedInvoices, invoices)
}

func TestGetInvoicesReturnsError(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := getInvoicesResponse{
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
		Get("/getInvoices").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		JSON(expectedResponse)

	balance, err := c.GetInvoices(nil)
	require.Nil(t, balance)
	require.ErrorContains(t, err, "getInvoices request error: code - 401, name - Unauthorized")
}
