package goswap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (p *Provider) GetCurrency(symbol string) (models.GetCurrency, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		p.BuildURL("get_currency", "&symbol=", symbol),
		nil,
	)
	if err != nil {
		return models.GetCurrency{}, err
	}

	statusCode, bodyBytes, err := p.RequestDoBytes(req)
	if err != nil {
		return models.GetCurrency{}, err
	}

	if statusCode != 200 {
		apiError := &models.APIError{}

		if err := json.Unmarshal(bodyBytes, apiError); err != nil {
			return models.GetCurrency{}, err
		}

		return models.GetCurrency{}, apiError
	}

	var response models.GetCurrency

	err = json.Unmarshal(bodyBytes, &response)

	return response, err
}
