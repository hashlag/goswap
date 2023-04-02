package models

import "fmt"

type APIError struct {
	Status      int    `json:"status"`
	ErrorText   string `json:"error"`
	Description string `json:"description"`
	TraceID     string `json:"traceId"`
}

func (re *APIError) Error() string {
	return fmt.Sprintf("status %d: %s (%s); trace id: %s",
		re.Status,
		re.ErrorText,
		re.Description,
		re.TraceID,
	)
}

type GetCurrency struct {
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
