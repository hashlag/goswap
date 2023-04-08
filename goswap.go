package goswap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/wachttijd/goswap/models"
)

type Provider struct {
	APIBase   string
	APIKey    string
	UserAgent string
	Client    http.Client
}

func NewProvider(apiBase, apiKey string) *Provider {
	return &Provider{
		APIBase:   strings.TrimSuffix(apiBase, "/"),
		APIKey:    apiKey,
		UserAgent: "goswap",
		Client:    http.Client{},
	}
}

func (p *Provider) BuildURL(endpoint string, fragments ...string) string {
	return fmt.Sprintf("%s/%s?api_key=%s%s",
		p.APIBase,
		endpoint,
		p.APIKey,
		strings.Join(fragments, ""),
	)
}

func (p *Provider) RequestDoBytes(req *http.Request) (int, []byte, error) {
	resp, err := p.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}

	respBodyBytes, err := io.ReadAll(resp.Body)

	return resp.StatusCode, respBodyBytes, err
}

func (p *Provider) GetCurrency(symbol string) (models.Currency, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_currency", "&symbol=", symbol),
		nil,
	)
	if err != nil {
		return models.Currency{}, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return models.Currency{}, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return models.Currency{}, err
		}

		return models.Currency{}, apiError
	}

	var r models.Currency

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetAllCurrencies() ([]models.Currency, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_all_currencies"),
		nil,
	)
	if err != nil {
		return nil, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return nil, err
		}

		return nil, apiError
	}

	var r []models.Currency

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetPairs(fixed bool, symbol string) ([]string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_pairs", "&fixed=", strconv.FormatBool(fixed), "&symbol=", symbol),
		nil,
	)
	if err != nil {
		return nil, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return nil, err
		}

		return nil, apiError
	}

	var r []string

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetAllPairs(fixed bool) (map[string][]string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_all_pairs", "&fixed=", strconv.FormatBool(fixed)),
		nil,
	)
	if err != nil {
		return nil, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return nil, err
		}

		return nil, apiError
	}

	var r map[string][]string

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) CreateExchange(exchange models.ExchangeConfig, XForwardedFor, XUserAgent string) (models.Exchange, error) {
	reqBody, err := json.Marshal(exchange)
	if err != nil {
		return models.Exchange{}, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		p.BuildURL("create_exchange"),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return models.Exchange{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	if XForwardedFor != "" {
		req.Header.Set("X-Forwarded-For", XForwardedFor)
	}

	if XUserAgent != "" {
		req.Header.Set("X-User-Agent", XUserAgent)
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return models.Exchange{}, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return models.Exchange{}, err
		}

		return models.Exchange{}, apiError
	}

	var r models.Exchange

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetExchange(id string) (models.Exchange, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_exchange", "&id=", id),
		nil,
	)
	if err != nil {
		return models.Exchange{}, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return models.Exchange{}, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return models.Exchange{}, err
		}

		return models.Exchange{}, apiError
	}

	var r models.Exchange

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetExchanges(limit, offset int, gte, lte string) ([]models.Exchange, error) {
	var optionals string

	if gte != "" {
		optionals += "&gte=" + gte
	}

	if lte != "" {
		optionals += "&lte=" + lte
	}

	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL(
			"get_exchanges",
			"&limit=", strconv.Itoa(limit),
			"&offset=", strconv.Itoa(offset),
			optionals,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return nil, err
		}

		return nil, apiError
	}

	var r []models.Exchange

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetRanges(fixed bool, currencyFrom, currencyTo string) (models.Ranges, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL(
			"get_ranges",
			"&fixed=", strconv.FormatBool(fixed),
			"&currency_from=", currencyFrom,
			"&currency_to=", currencyTo,
		),
		nil,
	)
	if err != nil {
		return models.Ranges{}, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return models.Ranges{}, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return models.Ranges{}, err
		}

		return models.Ranges{}, apiError
	}

	var r models.Ranges

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetEstimated(fixed bool, currencyFrom, currencyTo string, amount float64) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL(
			"get_estimated",
			"&fixed=", strconv.FormatBool(fixed),
			"&currency_from=", currencyFrom,
			"&currency_to=", currencyTo,
			"&amount=", strconv.FormatFloat(amount, 'f', -1, 64),
		),
		nil,
	)
	if err != nil {
		return "", err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return "", err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return "", err
		}

		return "", apiError
	}

	var r string

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) CheckExchanges(fixed bool, currencyFrom, currencyTo string, amount float64) (bool, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL(
			"check_exchanges",
			"&fixed=", strconv.FormatBool(fixed),
			"&currency_from=", currencyFrom,
			"&currency_to=", currencyTo,
			"&amount=", strconv.FormatFloat(amount, 'f', -1, 64),
		),
		nil,
	)
	if err != nil {
		return false, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return false, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return false, err
		}

		return false, apiError
	}

	var r bool

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetMarketInfo() ([]models.MarketInfo, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_market_info"),
		nil,
	)
	if err != nil {
		return nil, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return nil, err
		}

		return nil, apiError
	}

	var r []models.MarketInfo

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}
