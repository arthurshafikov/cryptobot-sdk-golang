package main

import (
	"log"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func getInvoices(client *cryptobot.Client) {
	invoices, err := client.GetInvoices(&cryptobot.GetInvoicesRequest{
		Asset:      "",
		InvoiceIDs: "",
		Status:     "",
		Offset:     0,
		Count:      100,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, invoice := range invoices {
		showInvoiceInfo(&invoice)
	}
}
