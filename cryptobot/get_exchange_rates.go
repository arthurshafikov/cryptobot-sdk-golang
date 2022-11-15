package cryptobot

import "fmt"

type GetExchangeRatesResponse struct {
	Response
	Result ExchangeRates `json:"result"`
}

type ExchangeRates []ExchangeRate

// for example
// cryptobot.ExchangeRate{IsValid:true, Rate:"16836.58000000", Source:"BTC", Target:"USD"}
type ExchangeRate struct {
	IsValid bool   `json:"is_valid"`
	Rate    string `json:"rate"`
	Source  string `json:"source"`
	Target  string `json:"target"`
}

// Use this method to get exchange rates of supported currencies. Returns slice of exchange rates.
func (c *Client) GetExchangeRates() (ExchangeRates, error) {
	responseBodyReader, err := c.request("getExchangeRates", nil)
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response GetExchangeRatesResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return response.Result, nil
	} else {
		return nil, fmt.Errorf("getExchangeRates request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
