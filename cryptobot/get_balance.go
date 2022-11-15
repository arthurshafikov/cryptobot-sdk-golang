package cryptobot

import "fmt"

type GetBalanceResponse struct {
	Response
	Result Balance `json:"result"`
}

type Balance []BalanceAsset

// for example cryptobot.BalanceAsset{Available:"0", CurrencyCode:"BTC"}
type BalanceAsset struct {
	Available    string `json:"available"`
	CurrencyCode string `json:"currency_code"`
}

// Use this method to get a balance of your app. Returns slice of BalanceAssets.
func (c *Client) GetBalance() (Balance, error) {
	responseBodyReader, err := c.request("getBalance", nil)
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response GetBalanceResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return response.Result, nil
	} else {
		return nil, fmt.Errorf("getBalance request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
