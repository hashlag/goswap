package goswap_test

import (
	"testing"

	"github.com/wachttijd/goswap"
)

func TestGetCurrency(t *testing.T) {
	provider := goswap.NewProvider(
		"https://api.simpleswap.io",
		"71c46760-92cf-45f4-b5c1-0bbf8031283b",
	)

	currency, err := provider.GetCurrency("btc")

	if err != nil {
		t.Error(err)
	}

	if currency.Name != "Bitcoin" || currency.Symbol != "btc" {
		t.Error("wrong currency name or/and symbol")
	}
}
