package cryptobot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseWebhookUpdate(t *testing.T) {
	requestDate := "2022-11-17T06:38:48.010Z"
	requestDateTime, err := time.Parse("2006-01-02T15:04:05Z07:00", requestDate)
	require.NoError(t, err)
	expectedWebhookUpdate := &WebhookUpdate{
		UpdateID:    123,
		UpdateType:  InvoicePaidWebhookUpdateType,
		RequestDate: requestDateTime,
	}

	data := []byte(`{"update_id":123,"update_type":"invoice_paid","request_date":"2022-11-17T06:38:48.010Z"}`)
	webhookUpdate, err := ParseWebhookUpdate(data)
	require.NoError(t, err)
	require.Equal(t, expectedWebhookUpdate, webhookUpdate)
}

func TestParseInvoice(t *testing.T) {
	data := []byte(`{"update_id":0,"update_type":"","request_date":"0001-01-01T00:00:00Z",
		"payload":{"invoice_id":123,"status":"paid","hash":"someHash","asset":"USDT","amount":"100",
		"fee":"2","pay_url":"some-pay-url","description":"some description","created_at":"2022-11-24T05:29:38.495Z",
		"usd_rate":"usdRate","allow_comments":false,"allow_anonymous":false,"expiration_date":"2022-11-30T05:29:38.495Z",
		"paid_at":"2022-11-28T05:29:38.495Z","paid_anonymously":false,"comment":"some comment","hidden_message":"some message",
		"payload":"some payload","paid_btn_name":"btn name","paid_btn_url":"btn url"}}`)

	invoice, err := ParseInvoice(data)
	require.NoError(t, err)
	require.Equal(t, expectedInvoice, *invoice)
}
