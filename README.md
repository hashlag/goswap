# GOSWAP — Third-party [simpleswap.io](https://simpleswap.io/?ref=d022a802baee) API library for Golang

This library allows you to easily integrate you Golang project with [simpleswap.io](https://simpleswap.io/?ref=d022a802baee).

You can get details about the API in the [documentation](https://api.simpleswap.io/) and the [affiliate program description](https://simpleswap.io/affiliate-program/how-to-start/api?ref=d022a802baee).

## Installation

```bash
go get github.com/hashlag/goswap
```

## Usage

Import GOSWAP, API models and create a provider:

```go
import (
	"github.com/hashlag/goswap"
	"github.com/hashlag/goswap/models"
)

var provider = goswap.NewProvider(
	"https://api.simpleswap.io",
	"API-KEY",
)
```

### [Get currency info](https://api.simpleswap.io/#/Currency/CurrencyController_getCurrency) and error handling example

```go
currency, err := provider.GetCurrency("btc")

if err != nil {
    if errors.Is(err, &models.APIError{}) {
    	// it's an API error
    	// you can retrieve some details about the error and handle it properly
    	fmt.Println(err.(*models.APIError).ErrorText)
    	fmt.Println(err.(*models.APIError).Description)
    	// etc.
    } else {
        // it is not an API error
    }
    // ...
}

fmt.Println(currency.Name) // Bitcoin
```

`provider.GetCurrency` takes 1 required argument:

- currency symbol — `string`

`currency` is a `models.Currency`:

```go
type Currency struct {
	Name              string   `json:"name"`
	Symbol            string   `json:"symbol"`
	Network           string   `json:"network"`
	HasExtraID        bool     `json:"has_extra_id"`
	ExtraID           string   `json:"extra_id"`
	Image             string   `json:"image"`
	WarningsFrom      []string `json:"warnings_from"`
	WarningsTo        []string `json:"warnings_to"`
	ValidationAddress string   `json:"validation_address"`
	ValidationExtra   string   `json:"validation_extra"`
	AddressExplorer   string   `json:"address_explorer"`
	TxExplorer        string   `json:"tx_explorer"`
	ConfirmationsFrom string   `json:"confirmations_from"`
}
```

### [Get all currencies](https://api.simpleswap.io/#/Currency/CurrencyController_getAllCurrencies)

```go
currencies, err := provider.GetAllCurrencies()
if err != nil {
    // ...
}
```

`currencies` is a `[]models.Currency`

### [Get exchnage pairs for currency](https://api.simpleswap.io/#/Pairs/PairController_getPair)

```go
pairs, err := provider.GetPairs(false, "btc")
if err != nil {
    // ...
}
```

`provider.GetPairs` takes 2 required arguments:

- fixed rate — `bool`
- currency symbol — `string`

`pairs` is a `[]string` containing allowed currency pairs symbols

### [Get all exchange pairs](https://api.simpleswap.io/#/Pairs/PairController_getAllPairs)

```go
pairs, err := provider.GetAllPairs(false)
if err != nil {
	// ...
}
```

`provider.GetAllPairs` takes 1 required argument:

- fixed rate — `bool`

`pairs` is a `map[string][]string` where each key is a currency and value is a list of allowed pairs to the currency

### [Create new exchnage](https://api.simpleswap.io/#/Exchange/ExchangeController_createExchange)

```go
exchange, err := provider.CreateExchange(
	models.ExchangeConfig{
		Fixed:             false,
		CurrencyFrom:      "ltc",
		CurrencyTo:        "btc",
		Amount:            1,
		AddressTo:         "1KyG8cnqU6TpZbHBMUCcsARxKLmzPEhf9",
		ExtraIDTo:         "",
		UserRefundAddress: "LXMt7yYkPvsdpXa29By38p4VYnurpECnrs",
		UserRefundExtraID: "",
	},
	"", // XForwardedFor is optional
	"", // XUserAgent is optional too...
)
if err != nil {
	// ...
}

fmt.Println(exchange.Status) // "waiting", for example
```

`provider.CreateExchange` takes 1 required argument:

- exchnage config — `models.ExchangeConfig`

And 2 arguments that can be omitted:

- X-Forwarded-For header — `string`
- X-User-Agent header — `string`

`models.ExchangeConfig` is:

```go
type ExchangeConfig struct {
	// Pass "true" if you want to request fixed-rate pairs, any other value for floating-rate pairs
	Fixed             bool    `json:"fixed"`
	// The ticker of a currency the client will transfer
	CurrencyFrom      string  `json:"currency_from"`
	// The ticker of a currency the client will receive
	CurrencyTo        string  `json:"currency_to"`
	// The amount which the client is willing to transfer
	Amount            float64 `json:"amount"`
	// Address where the client will receive payment
	AddressTo         string  `json:"address_to"`
	// Memo/address ID to to obtain the exchange result (optional, not all networks require it)
	ExtraIDTo         string  `json:"extra_id_to"`
	// Address where the client will get a refund (optional)
	UserRefundAddress string  `json:"user_refund_address"`
	// Memo/address ID to get a refund (optional, not all networks require it)
	UserRefundExtraID string  `json:"user_refund_extra_id"`
}
```

`exchange` is a `models.Exchange`:

```go
type Exchange struct {
	ID                string              `json:"id"`
	Type              string              `json:"type"`
	Timestamp         string              `json:"timestamp"`
	UpdatedAt         string              `json:"updated_at"`
	CurrencyFrom      string              `json:"currency_from"`
	CurrencyTo        string              `json:"currency_to"`
	AmountFrom        string              `json:"amount_from"`
	ExpectedAmount    string              `json:"expected_amount"`
	AmountTo          string              `json:"amount_to"`
	AddressFrom       string              `json:"address_from"`
	AddressTo         string              `json:"address_to"`
	ExtraIDFrom       string              `json:"extra_id_from"`
	ExtraIDTo         string              `json:"extra_id_to"`
	UserRefundAddress string              `json:"user_refund_address"`
	UserRefundExtraID string              `json:"user_refund_extra_id"`
	TxFrom            string              `json:"tx_from"`
	TxTo              string              `json:"tx_to"`
	Status            string              `json:"status"`
	RedirectURL       string              `json:"redirect_url"`
	Currencies        map[string]Currency `json:"currencies"`
}
```

See an [exchange object schema](https://api.simpleswap.io/#/Exchange/ExchangeController_getExchange) for details.

### [Get info about exchange](https://api.simpleswap.io/#/Exchange/ExchangeController_getExchange)

```go
exchange, err := provider.GetExchange("EXCHANGE-ID")
if err != nil {
	// ...
}

fmt.Println(exchange.Status) // "confirming", for example
```

`provider.GetExchange` takes 1 required argument:

- exchange id — `string`

`exchange` is a `models.Exchange`

### [Get info about created exchanges](https://api.simpleswap.io/#/Exchange/ExchangeController_getAllExchanges)

```go
exchanges, err := provider.GetExchanges(50, 0, "", "")
if err != nil {
	// ...
}

fmt.Println(exchanges[0].Status) // "finished", for example
```

`provider.GetExchanges` takes 2 required arguments:

- limit of transactions to return (bigger than 0 and less than 1000) — `int`
- number of transactions to skip (bigger than 0) — `int`

And 2 arguments that can be omitted:

- minimum display date, format must be ISO 8601. ('2000-01-01' or '2000-01-01T00:00:00Z') — `string`
- maximum date to display, format must be ISO 8601. ('2000-01-01' or '2000-01-01T00:00:00Z') — `string`

`exchanges` is a `[]models.Exchange`

### [Get minimal and maximal (if exists) amount for exchange between selected currencies](https://api.simpleswap.io/#/Exchange/ExchangeController_getRange)

```go
ranges, err := provider.GetRanges(false, "ltc", "btc")
if err != nil {
	// ...
}

fmt.Printf("min: %s max: %s\n", ranges.Min, ranges.Max)
```

`provider.GetRanges` takes 3 required arguments:

- fixed rate — `bool`
- the ticker of a currency the client will transfer — `string`
- the ticker of a currency the client will receive — `string`

`ranges` is a `models.Ranges`:

```go
type Ranges struct {
	Min string `json:"min"`
	Max string `json:"max"`
}
```

### [Get estimated exchange amount](https://api.simpleswap.io/#/Exchange/ExchangeController_getEstimated)

```go
estimated, err := provider.GetEstimated(false, "ltc", "btc", float64(1.77))
if err != nil {
	// ...
}

fmt.Println(estimated) // "0.0054897", for example
```

`provider.GetEstimated` takes 4 required arguments:

- fixed rate — `bool`
- the ticker of a currency the client will transfer — `string`
- the ticker of a currency the client will receive — `string`
- amount to exchange — `float64`

`estimated` is a `string`

### [Check exchange's parameters](https://api.simpleswap.io/#/Exchange/ExchangeController_checkExchanges)

```go
possible, err := provider.CheckExchanges(false, "ltc", "btc", float64(1.77))
if err != nil {
	// ...
}

fmt.Println(possible) // true, for example
```

`provider.CheckExchanges` takes 4 required arguments:

- fixed rate — `bool`
- the ticker of a currency the client will transfer — `string`
- the ticker of a currency the client will receive — `string`
- amount to exchange — `float64`

`possible` is a `bool`

### [Get full market info (only fixed rate)](https://api.simpleswap.io/#/Market/MarketController_getMarketInfo)

```go
marketInfo, err := provider.GetMarketInfo()
if err != nil {
	// ...
}

fmt.Printf("%+v\n", marketInfo[0])
```

`marketInfo` is a `[]models.MarketInfo`

`models.MarketInfo`:

```go
type MarketInfo struct {
	CurrencyFrom string `json:"currency_from"`
	CurrencyTo   string `json:"currency_to"`
	Min          string `json:"min"`
	Max          string `json:"max"`
	Rate         string `json:"rate"`
}
```

### Custom (API) errors

You can check if returned error is an API error with `errors.Is()`:

```go
if errors.Is(err, &models.APIError{}) {
    	// it's an API error
    	// you can retrieve some details about the error and handle it properly
    	fmt.Println(err.(*models.APIError).ErrorText)
    	fmt.Println(err.(*models.APIError).Description)
    	// etc.
} else {
	// it is not an API error
}
```

`models.APIError`:

```go
type APIError struct {
	Status      int    `json:"status"`
	ErrorText   string `json:"error"`
	Description string `json:"description"`
	TraceID     string `json:"traceId"`
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](https://raw.githubusercontent.com/hashlag/goswap/main/LICENSE) file for details.
