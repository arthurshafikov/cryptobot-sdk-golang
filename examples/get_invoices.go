package main

import (
	"log"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func getInvoices(client *cryptobot.Client) {
	// GetInvoicesRequest argument is completely optional here
	invoices, err := client.GetInvoices(&cryptobot.GetInvoicesRequest{
		Asset:      "USDT",
		InvoiceIDs: "1,2,3",
		Status:     "active",
		Offset:     0,
		Count:      20,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, invoice := range invoices {
		showInvoiceInfo(&invoice)
	}
}
