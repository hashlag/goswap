package models

type ExchangeConfig struct {
	Fixed             bool    `json:"fixed"`
	CurrencyFrom      string  `json:"currency_from"`
	CurrencyTo        string  `json:"currency_to"`
	Amount            float64 `json:"amount"`
	AddressTo         string  `json:"address_to"`
	ExtraIDTo         string  `json:"extra_id_to"`
	UserRefundAddress string  `json:"user_refund_address"`
	UserRefundExtraID string  `json:"user_refund_extra_id"`
}
