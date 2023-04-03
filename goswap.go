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

func (p *Provider) GetPairs(fixed bool, symbol string) (models.Pairs, error) {
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

	var r models.Pairs

	err = json.Unmarshal(bodyBytes, &r)

	return r, err
}

func (p *Provider) GetAllPairs(fixed bool) (models.AllPairs, error) {
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

	var r models.AllPairs

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
