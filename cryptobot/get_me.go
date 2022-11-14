package cryptobot

import (
	"encoding/json"
	"io/ioutil"
)

type GetMeResponse struct {
	Response
	Result AppInfo `json:"result"`
}

type AppInfo struct {
	AppID                        int64  `json:"app_id"`
	Name                         string `json:"name"`
	PaymentProcessingBotUsername string `json:"payment_processing_bot_username"`
}

func (c *Client) GetMe() (*AppInfo, error) {
	responseBodyReader, err := c.request("getMe", nil)
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response GetMeResponse
	responseBody, err := ioutil.ReadAll(responseBodyReader)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return &response.Result, nil
}
