package main

import (
	"log"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func parseWebhook() {
	// let's say we have received this request body
	requestBody := []byte(`{"update_id":123,"update_type":"invoice_paid","request_date":"2022-11-17T06:38:48.010Z", 
		"payload":{"invoice_id":123,"status":"paid","hash":"someHash","asset":"USDT","amount":"100",
		"fee":"2","pay_url":"some-pay-url","description":"some description","created_at":"2022-11-24T05:29:38.495Z",
		"usd_rate":"usdRate","allow_comments":false,"allow_anonymous":false,"expiration_date":"2022-11-30T05:29:38.495Z",
		"paid_at":"2022-11-28T05:29:38.495Z","paid_anonymously":false,"comment":"some comment","hidden_message":"some message",
		"payload":"some payload","paid_btn_name":"btn name","paid_btn_url":"btn url"}}`)

	webhookUpdate, err := cryptobot.ParseWebhookUpdate(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	switch webhookUpdate.UpdateType {
	case cryptobot.InvoicePaidWebhookUpdateType:
		invoice, err := cryptobot.ParseInvoice(requestBody)
		if err != nil {
			log.Fatalln(err)
		}

		if invoice.Status == cryptobot.InvoiceActiveStatus {
			showInvoiceInfo(invoice)
		}
	default:
		log.Fatalln("unsupported webhook update type " + webhookUpdate.UpdateType)
	}
}
