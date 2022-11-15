package main

import (
	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func main() {
	client := cryptobot.NewClient(cryptobot.Options{
		Testing:  true,
		APIToken: "<YOUR_API_KEY>",
	})

	createInvoice(client)
	getBalance(client)
	getCurrencies(client)
	getExchangeRates(client)
	getMe(client)

	// these methods are disabled in test mode
	transfer(client)
	getInvoices(client)
}
