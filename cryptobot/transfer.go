package cryptobot

import (
	"fmt"
	"net/url"
	"strconv"
)

type TransferRequest struct {
	// Telegram user ID. User must have previously used @CryptoBot (@CryptoTestnetBot for testnet).
	UserID int64 `json:"user_id"`

	// Currency code. Currently, can be “BTC”, “TON”, “ETH”, “USDT”, “USDC” or “BUSD”.
	Asset string `json:"asset"`

	// Amount of the transfer in float. The minimum and maximum amounts for each of
	// the support asset roughly correspond to the limit of 1-25000 USD. Use getExchangeRates to convert amounts.
	// For example: 125.50
	Amount string `json:"amount"`

	// Unique ID to make your request idempotent and ensure that only one of the transfers with the same spend_id is accepted from your app.
	// This parameter is useful when the transfer should be retried (i.e. request timeout, connection reset, 500 HTTP status, etc).
	// Up to 64 symbols.
	SpendID string `json:"spend_id"`

	// Optional. Comment for the transfer. Users will see this comment when they receive a notification about the transfer.
	// Up to 1024 symbols.
	Comment string `json:"comment"`

	// Optional. Pass true if the user should not receive a notification about the transfer. Default is false.
	DisableSendNotification bool `json:"disable_send_notification"`
}

type TransferResponse struct {
	Response
	Result Transfer `json:"result"`
}

// Use this method to send coins from your app's balance to a user. On success, returns struct of completed transfer.
func (c *Client) Transfer(transferRequest *TransferRequest) (*Transfer, error) {
	responseBodyReader, err := c.request("transfer", func(q url.Values) url.Values {
		q.Add("user_id", strconv.FormatInt(transferRequest.UserID, 10))
		q.Add("asset", transferRequest.Asset)
		q.Add("amount", transferRequest.Amount)
		q.Add("spend_id", transferRequest.SpendID)
		if transferRequest.Comment != "" {
			q.Add("comment", transferRequest.Comment)
		}
		// default is false
		if transferRequest.DisableSendNotification {
			q.Add("allow_anonymous", "true")
		}

		return q
	})
	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response TransferResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return &response.Result, nil
	} else {
		return nil, fmt.Errorf("transfer request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
	}
}
