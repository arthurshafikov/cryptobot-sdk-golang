package cryptobot

import (
	"fmt"
	"net/url"
	"strconv"
)

type GetInvoicesRequest struct {
	// Optional. Currency codes separated by comma.
	// Supported assets: “BTC”, “TON”, “ETH”, “USDT”, “USDC” and “BUSD”. Defaults to all assets.
	Asset string `json:"asset"`

	// Optional. Invoice IDs separated by comma.
	InvoiceIDs string `json:"invoice_ids"`

	// Optional. Status of invoices to be returned. Available statuses: “active” and “paid”. Defaults to all statuses.
	Status string `json:"status"`

	// Optional. Offset needed to return a specific subset of invoices. Default is 0.
	Offset int `json:"offset"`

	// Optional. Number of invoices to be returned. Values between 1-1000 are accepted. Default is 100.
	Count int `json:"count"`
}

type GetInvoicesResponse struct {
	Response
	Result Invoices `json:"result"`
}

type Invoices struct {
	Items []Invoice `json:"items"`
}

// Use this method to get invoices of your app. On success, returns slice of invoices.
func (c *Client) GetInvoices(getInvoicesRequest *GetInvoicesRequest) ([]Invoice, error) {
	responseBodyReader, err := c.request("getInvoices", func(q url.Values) url.Values {
		if getInvoicesRequest.Asset != "" {
			q.Add("asset", getInvoicesRequest.Asset)
		}
		if getInvoicesRequest.InvoiceIDs != "" {
			q.Add("invoice_ids", getInvoicesRequest.InvoiceIDs)
		}
		if getInvoicesRequest.Status != "" {
			q.Add("status", getInvoicesRequest.Status)
		}
		if getInvoicesRequest.Offset != 0 {
			q.Add("offset", strconv.Itoa(getInvoicesRequest.Offset))
		}
		if getInvoicesRequest.Count != 0 {
			q.Add("count", strconv.Itoa(getInvoicesRequest.Offset))
		}

		return q
	})
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response GetInvoicesResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", response)

	if response.Ok {
		return response.Result.Items, nil
	} else {
		return nil, fmt.Errorf("getInvoices request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
