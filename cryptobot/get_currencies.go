package cryptobot

import "fmt"

type getCurrenciesResponse struct {
	response
	Result Currencies `json:"result"`
}

type Currencies []Currency

// for example
// cryptobot.Currency{Code:"BTC", Decimals:8, IsBlockchain:true, IsFiat:false, IsStablecoin:false, Name:"Bitcoin", Url:"https://bitcoin.org/"}
// cryptobot.Currency{Code:"USD", Decimals:8, IsBlockchain:false, IsFiat:true, IsStablecoin:false, Name:"United States dollar", Url:""}
type Currency struct {
	Code         string `json:"code"`
	Decimals     int    `json:"decimals"`
	IsBlockchain bool   `json:"is_blockchain"`
	IsFiat       bool   `json:"is_fiat"`
	IsStablecoin bool   `json:"is_stablecoin"`
	Name         string `json:"name"`
	Url          string `json:"url"`
}

// Use this method to get a list of supported currencies. Returns slice of currencies.
func (c *Client) GetCurrencies() (Currencies, error) {
	responseBodyReader, err := c.request("getCurrencies", nil)
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response getCurrenciesResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return response.Result, nil
	} else {
		return nil, fmt.Errorf("getCurrencies request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
