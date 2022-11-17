# Cryptobot SDK Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/arthurshafikov/cryptobot-sdk-golang.svg)](https://pkg.go.dev/github.com/arthurshafikov/cryptobot-sdk-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurshafikov/cryptobot-sdk-golang)](https://goreportcard.com/report/github.com/arthurshafikov/cryptobot-sdk-golang)
![Tests](https://github.com/arthurshafikov/cryptobot-sdk-golang/actions/workflows/tests.yml/badge.svg)
![License](https://img.shields.io/github/license/arthurshafikov/cryptobot-sdk-golang)

Convenient SDK for [@Cryptobot](https://t.me/CryptoBot)

Use it if you would like to accept payments in cryptocurrency through your Telegram Bot with Golang.

# Installation

```
go get github.com/arthurshafikov/cryptobot-sdk-golang
```


# Quick Start

```golang
import "github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"

func main() {
	client := cryptobot.NewClient(cryptobot.Options{
		APIToken: "<YOUR_API_KEY>",
	})

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
```

# Examples

## [CreateInvoiceExample](examples/create_invoice.go)

```golang
invoice, err := client.CreateInvoice(cryptobot.CreateInvoiceRequest{
    Asset:          "USDT",
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
```

## [GetBalanceExample](examples/get_balance.go)

```golang
balance, err := client.GetBalance()
if err != nil {
    log.Fatalln(err)
}

for _, asset := range balance {
    fmt.Printf("Currency - %s, available - %s\n", asset.CurrencyCode, asset.Available)
}
```

## [GetCurrenciesExample](examples/get_currencies.go)

```golang
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
```

## [GetExchangeRatesExample](examples/get_exchange_rates.go)

```golang
exchangeRates, err := client.GetExchangeRates()
if err != nil {
    log.Fatalln(err)
}

for _, exchangeRate := range exchangeRates {
    fmt.Printf(
        "IsValid - %v, Rate - %v, Source - %v, Target - %v\n",
        exchangeRate.IsValid,
        exchangeRate.Rate,
        exchangeRate.Source,
        exchangeRate.Target,
    )
}
```

## [GetInvoicesExample](examples/get_invoices.go)

```golang
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
```

## [TransferExample](examples/transfer.go) 

```golang
transfer, err := client.Transfer(cryptobot.TransferRequest{
    UserID:                  1,
    Asset:                   "TON",
    Amount:                  "10.5",
    SpendID:                 "",
    Comment:                 "Debt",
    DisableSendNotification: false,
})
if err != nil {
    log.Fatalln(err)
}

fmt.Printf(
    "ID - %v, UserID - %s, Status - %s, Amount - %s, Asset - %s, Comment - %s, CompletedAt - %s \n",
    transfer.ID,
    transfer.UserID,
    transfer.Status,
    transfer.Amount,
    transfer.Asset,
    transfer.Comment,
    transfer.CompletedAt,
)
```

## [WebhookExample](examples/webhook.go) 

```golang
requestBody := []byte(`someJSONRequestBodyContent`)

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

    showInvoiceInfo(invoice)
default:
    log.Fatalln("unsupported webhook update type " + webhookUpdate.UpdateType)
}
```

# Documentation

Check out this repository [documentation](https://pkg.go.dev/github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot)

Check out official CryptoBot [documentation](https://help.crypt.bot/crypto-pay-api)

# Troubleshooting

Feel free to create Issues and even suggest your own PRs
