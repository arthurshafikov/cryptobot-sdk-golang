package main

import (
	"fmt"
	"log"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func getCurrencies(client *cryptobot.Client) {
	currencies, err := client.GetCurrencies()
	if err != nil {
		log.Fatalln(err)
	}

	for _, currency := range currencies {
		fmt.Printf(
			"Code - %s, Decimals - %v, IsBlockchain - %v, IsFiat - %v, IsStablecoin - %v, Name - %s, Url - %s \n",
			currency.Code,
			currency.Decimals,
			currency.IsBlockchain,
			currency.IsFiat,
			currency.IsStablecoin,
			currency.Name,
			currency.Url,
		)
	}
}
