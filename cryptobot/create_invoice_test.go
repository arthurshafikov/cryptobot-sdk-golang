package cryptobot

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

var expectedInvoice = Invoice{
	ID:              123,
	Status:          "paid",
	Hash:            "someHash",
	Asset:           "USDT",
	Amount:          "100",
	Fee:             "2",
	PayUrl:          "some-pay-url",
	Description:     "some description",
	CreatedAt:       "2022-11-24T05:29:38.495Z",
	UsdRate:         "usdRate",
	AllowComments:   false,
	AllowAnonymous:  false,
	ExpirationDate:  "2022-11-30T05:29:38.495Z",
	PaidAt:          "2022-11-28T05:29:38.495Z",
	PaidAnonymously: false,
	Comment:         "some comment",
	HiddenMessage:   "some message",
	Payload:         "some payload",
	PaidBtnName:     "btn name",
	PaidBtnUrl:      "btn url",
}

func TestCreateInvoice(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := createInvoiceResponse{
		response: response{
			Ok: true,
		},
		Result: expectedInvoice,
	}
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/createInvoice").
		MatchHeader(apiTokenHeaderName, testToken).
		MatchParam("asset", "USDT").
		MatchParam("amount", "100").
		MatchParam("description", "some description").
		MatchParam("hidden_message", "some message").
		MatchParam("paid_btn_name", "btn name").
		MatchParam("paid_btn_url", "btn url").
		MatchParam("payload", "some payload").
		MatchParam("allow_comments", "false").
		MatchParam("allow_anonymous", "false").
		MatchParam("expires_in", "123123").
		Reply(200).
		JSON(expectedResponse)

	invoice, err := c.CreateInvoice(CreateInvoiceRequest{
		Asset:          "USDT",
		Amount:         "100",
		Description:    "some description",
		HiddenMessage:  "some message",
		PaidBtnName:    "btn name",
		PaidBtnUrl:     "btn url",
		Payload:        "some payload",
		AllowComments:  false,
		AllowAnonymous: false,
		ExpiresIn:      123123,
	})
	require.NoError(t, err)
	require.Equal(t, expectedInvoice, *invoice)
}

func TestCreateInvoiceReturnsError(t *testing.T) {
	c := getTestClient(t)
	expectedResponse := createInvoiceResponse{
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
		Get("/createInvoice").
		MatchHeader(apiTokenHeaderName, testToken).
		MatchParam("asset", "USDT").
		MatchParam("amount", "100").
		MatchParam("description", "some description").
		MatchParam("hidden_message", "some message").
		MatchParam("paid_btn_name", "btn name").
		MatchParam("paid_btn_url", "btn url").
		MatchParam("payload", "some payload").
		MatchParam("expires_in", "123123").
		Reply(200).
		JSON(expectedResponse)

	invoice, err := c.CreateInvoice(CreateInvoiceRequest{
		Asset:          "USDT",
		Amount:         "100",
		Description:    "some description",
		HiddenMessage:  "some message",
		PaidBtnName:    "btn name",
		PaidBtnUrl:     "btn url",
		Payload:        "some payload",
		AllowComments:  true,
		AllowAnonymous: true,
		ExpiresIn:      123123,
	})
	require.Nil(t, invoice)
	require.ErrorContains(t, err, "createInvoice request error: code - 403, name - Forbidden")
}
