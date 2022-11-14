package cryptobot

import (
	"fmt"
	"net/url"
	"strconv"
)

type CreateInvoiceRequest struct {
	// Currency code. Currently, can be “BTC”, “TON”, “ETH”, “USDT”, “USDC” or “BUSD”.
	Asset string `json:"asset"`

	// Amount of the invoice in float. For example: 125.50
	Amount string `json:"amount"`

	// Optional. Description for the invoice. User will see this description when they pay the invoice. Up to 1024 characters.
	Description string `json:"description"`

	// Optional. Text of the message that will be shown to a user after the invoice is paid. Up to 2048 characters.
	HiddenMessage string `json:"hidden_message"`

	// Optional. Name of the button that will be shown to a user after the invoice is paid.
	// Supported names:
	// viewItem – “View Item”
	// openChannel – “View Channel”
	// openBot – “Open Bot”
	// callback – “Return”
	PaidBtnName string `json:"paid_btn_name"`

	// Optional. Required if paid_btn_name is used.URL to be opened when the button is pressed.
	// You can set any success link (for example, a link to your bot). Starts with https or http.
	PaidBtnUrl string `json:"paid_btn_url"`

	// Optional. Any data you want to attach to the invoice (for example, user ID, payment ID, ect). Up to 4kb.
	Payload string `json:"payload"`

	// Optional. Allow a user to add a comment to the payment. Default is true.
	AllowComments bool `json:"allow_comments"`

	// Optional. Allow a user to pay the invoice anonymously. Default is true.
	AllowAnonymous bool `json:"allow_anonymous"`

	// Optional. You can set a payment time limit for the invoice in seconds. Values between 1-2678400 are accepted.
	ExpiresIn int64 `json:"expires_in"`
}

type CreateInvoiceResponse struct {
	Response
	Result Invoice `json:"result"`
}

// Use this method to create a new invoice. On success, returns an object of the created invoice.
func (c *Client) CreateInvoice(createInvoiceRequest *CreateInvoiceRequest) (*Invoice, error) {
	responseBodyReader, err := c.request("createInvoice", func(q url.Values) url.Values {
		q.Add("asset", createInvoiceRequest.Asset)
		q.Add("amount", createInvoiceRequest.Amount)
		if createInvoiceRequest.Description != "" {
			q.Add("description", createInvoiceRequest.Description)
		}
		if createInvoiceRequest.HiddenMessage != "" {
			q.Add("hidden_message", createInvoiceRequest.HiddenMessage)
		}
		if createInvoiceRequest.PaidBtnName != "" {
			q.Add("paid_btn_name", createInvoiceRequest.PaidBtnName)
		}
		if createInvoiceRequest.PaidBtnUrl != "" {
			q.Add("paid_btn_url", createInvoiceRequest.PaidBtnUrl)
		}
		if createInvoiceRequest.Payload != "" {
			q.Add("payload", createInvoiceRequest.Payload)
		}
		// default is true
		if !createInvoiceRequest.AllowComments {
			q.Add("allow_comments", "false")
		}
		// default is true
		if !createInvoiceRequest.AllowAnonymous {
			q.Add("allow_anonymous", "false")
		}
		if createInvoiceRequest.ExpiresIn != 0 {
			q.Add("expires_in", strconv.FormatInt(createInvoiceRequest.ExpiresIn, 10))
		}

		return q
	})
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response CreateInvoiceResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return &response.Result, nil
	} else {
		return nil, fmt.Errorf("createInvoice request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
