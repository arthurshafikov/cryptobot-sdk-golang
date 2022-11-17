package cryptobot

import (
	"encoding/json"
	"time"
)

const InvoicePaidWebhookUpdateType = "invoice_paid"

// Default WebhookUpdate that you will receive, always check the UpdateType field
// to know what to do with the request afterwards.
type WebhookUpdate struct {
	// Non-unique update ID.
	UpdateID int `json:"update_id"`

	// Webhook update type. Supported update types:
	// invoice_paid – the update sent when the invoice is paid.
	UpdateType string `json:"update_type"`

	// Date the request was sent in ISO 8601 format.
	RequestDate time.Time `json:"request_date"`

	// Payload may contain anything.
	Payload any `json:"payload"`
}

type InvoicePaidWebhookUpdate struct {
	WebhookUpdate

	// Payload contains Invoice object.
	Payload Invoice `json:"payload"`
}

// Convenient function-helper to parse WebhookUpdate from request body easily.
func ParseWebhookUpdate(data []byte) (*WebhookUpdate, error) {
	var webhookUpdate WebhookUpdate
	if err := json.Unmarshal(data, &webhookUpdate); err != nil {
		return nil, err
	}

	return &webhookUpdate, nil
}

// Convenient function-helper to parse Invoice from request body easily
// use it only if you're sure that UpdateType is invoice_paid.
func ParseInvoice(data []byte) (*Invoice, error) {
	var invoicePaidWebhookUpdate InvoicePaidWebhookUpdate
	if err := json.Unmarshal(data, &invoicePaidWebhookUpdate); err != nil {
		return nil, err
	}

	return &invoicePaidWebhookUpdate.Payload, nil
}
