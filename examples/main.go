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
	getInvoices(client)

	// these methods are disabled by default, check your app's security settings in @CryptoBot
	transfer(client)
}
