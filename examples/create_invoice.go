package main

import (
	"fmt"
	"log"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func createInvoice(client *cryptobot.Client) {
	invoice, err := client.CreateInvoice(cryptobot.CreateInvoiceRequest{
		Asset:          cryptobot.USDT,
		Amount:         "125.50",
		Description:    "Description for the user",
		HiddenMessage:  "After invoice is paid user will see this message",
		PaidBtnName:    "", // optional. one of these viewItem, openChannel, openBot, callback
		PaidBtnUrl:     "", // URL to be opened when the PaidBtn is pressed
		Payload:        "any payload we need in out application",
		AllowComments:  false,  // Allow a user to add a comment to the payment. Default is true
		AllowAnonymous: false,  // Allow a user to pay the invoice anonymously. Default is true.
		ExpiresIn:      60 * 5, // invoice will expire in 5 minutes
	})
	if err != nil {
		log.Fatalln(err)
	}

	showInvoiceInfo(invoice)
}

func showInvoiceInfo(invoice *cryptobot.Invoice) {
	fmt.Printf("Invoice ID: %v\n"+
		"Status: %s\n"+
		"Hash: %s\n"+
		"Asset: %s\n"+
		"Amount: %s\n"+
		"Fee: %s\n"+
		"PayUrl: %s\n"+
		"Description: %s\n"+
		"CreatedAt: %s\n"+
		"UsdRate: %s\n"+
		"AllowComments: %v\n"+
		"AllowAnonymous: %v\n"+
		"ExpirationDate: %s\n"+
		"PaidAt: %s\n"+
		"PaidAnonymously: %v\n"+
		"Comment: %s\n"+
		"HiddenMessage: %s\n"+
		"Payload: %s\n"+
		"PaidBtnName: %s\n"+
		"PaidBtnUrl: %s\n",
		invoice.ID,
		invoice.Status,
		invoice.Hash,
		invoice.Asset,
		invoice.Amount,
		invoice.Fee,
		invoice.PayUrl,
		invoice.Description,
		invoice.CreatedAt,
		invoice.UsdRate,
		invoice.AllowComments,
		invoice.AllowAnonymous,
		invoice.ExpirationDate,
		invoice.PaidAt,
		invoice.PaidAnonymously,
		invoice.Comment,
		invoice.HiddenMessage,
		invoice.Payload,
		invoice.PaidBtnName,
		invoice.PaidBtnUrl,
	)
}
