package goswap_test

import (
	"log"
	"testing"

	"github.com/wachttijd/goswap"
	"github.com/wachttijd/goswap/models"
)

const API_KEY = "API_KEY"

var provider = goswap.NewProvider(
	"https://api.simpleswap.io",
	API_KEY,
)

func TestKey(t *testing.T) {
	if API_KEY == "API_KEY" {
		t.Fail()
		log.Fatalln("insert your API key")
	}
}

func TestGetCurrency(t *testing.T) {
	currency, err := provider.GetCurrency("btc")

	if err != nil {
		// if errors.Is(err, &models.APIError{}) {
		// 	// it's an API error
		// 	// you can retrieve some details about it
		// 	t.Log(err.(*models.APIError).ErrorText)
		// 	t.Log(err.(*models.APIError).Description)
		// 	// etc.
		// }
		t.Error(err)
		return
	}

	if currency.Name != "Bitcoin" || currency.Symbol != "btc" {
		t.Error("wrong currency name or/and symbol")
	}
}

func TestGetAllCurrencies(t *testing.T) {
	currencies, err := provider.GetAllCurrencies()
	if err != nil {
		t.Error(err)
		return
	}

	if len(currencies) == 0 {
		t.Error("empty result")
	}
}

func TestGetPairs(t *testing.T) {
	fixedRatePairs, err := provider.GetPairs(false, "btc")
	if err != nil {
		t.Error(err)
		return
	}

	if len(fixedRatePairs) == 0 {
		t.Error("empty result (fixed-rate pairs)")
		return
	}

	floatingRatePairs, err := provider.GetPairs(true, "btc")
	if err != nil {
		t.Error(err)
		return
	}

	if len(floatingRatePairs) == 0 {
		t.Error("empty result (floating-rate pairs)")
	}
}

func TestGetAllPairs(t *testing.T) {
	fixedRatePairs, err := provider.GetAllPairs(false)
	if err != nil {
		t.Error(err)
		return
	}

	if len(fixedRatePairs) == 0 {
		t.Error("empty result (fixed-rate pairs)")
		return
	}

	floatingRatePairs, err := provider.GetAllPairs(true)
	if err != nil {
		t.Error(err)
		return
	}

	if len(floatingRatePairs) == 0 {
		t.Error("empty result (floating-rate pairs)")
	}
}

func TestCreateExchange(t *testing.T) {
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
		t.Error(err)
		return
	}

	if exchange.Status != "waiting" {
		t.Error("exchange status is not 'waiting'")
	}
}

func TestGetExchange(t *testing.T) {
	// example exchange id
	exchange, err := provider.GetExchange("gnt3fwnw75t8796h")
	if err != nil {
		t.Error(err)
		return
	}

	if exchange.ID == "" || exchange.Timestamp == "" {
		t.Error("empty exchange id or/and timestamp")
	}
}

func TestGetExchanges(t *testing.T) {
	exchanges, err := provider.GetExchanges(50, 0, "", "")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("exchanges found: %d", len(exchanges))
}

func TestGetRanges(t *testing.T) {
	ranges, err := provider.GetRanges(false, "ltc", "btc")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("min: %s max: %s", ranges.Min, ranges.Max)
}

func TestGetEstimated(t *testing.T) {
	estimated, err := provider.GetEstimated(false, "ltc", "btc", float64(1.77))
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("estimated amount: %s", estimated)
}

func TestCheckExchnages(t *testing.T) {
	possible, err := provider.CheckExchanges(false, "ltc", "btc", float64(1.77))
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("possible to create an exchnage: %t", possible)
}

func TestGetMarketInfo(t *testing.T) {
	marketInfo, err := provider.GetMarketInfo()
	if err != nil {
		t.Error(err)
		return
	}

	if len(marketInfo) == 0 {
		t.Error("empty market info array")
	}
}
