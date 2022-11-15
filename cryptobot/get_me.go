package cryptobot

import "fmt"

type getMeResponse struct {
	response
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

	var response getMeResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return &response.Result, nil
	} else {
		return nil, fmt.Errorf("getMe request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
