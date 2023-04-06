package models

import "fmt"

type APIError struct {
	Status      int    `json:"status"`
	ErrorText   string `json:"error"`
	Description string `json:"description"`
	TraceID     string `json:"traceId"`
}

func (re *APIError) Error() string {
	return fmt.Sprintf("API error: status %d, %s (%s); trace id: %s",
		re.Status,
		re.ErrorText,
		re.Description,
		re.TraceID,
	)
}

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

type Pairs []string

type AllPairs map[string][]string

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

type Ranges struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type MarketInfo struct {
	CurrencyFrom string `json:"currency_from"`
	CurrencyTo   string `json:"currency_to"`
	Min          string `json:"min"`
	Max          string `json:"max"`
	Rate         string `json:"rate"`
}