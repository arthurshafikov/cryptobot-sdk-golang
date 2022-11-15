package main

import (
	"fmt"
	"log"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
)

func getMe(client *cryptobot.Client) {
	appInfo, err := client.GetMe()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(
		"AppID - %v, Name - %s, PaymentProcessingBotUsername - %s \n",
		appInfo.AppID,
		appInfo.Name,
		appInfo.PaymentProcessingBotUsername,
	)
}
