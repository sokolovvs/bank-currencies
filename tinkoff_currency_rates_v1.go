package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SuccessResponseFromTinkoffCurrencyRates struct {
	ResultCode string `json:"resultCode"`
	Payload    struct {
		LastUpdate struct {
			Milliseconds int `json:"milliseconds"`
		} `json:"lastUpdate"`
		Rates []RateFromResponse `json:"rates"`
	} `json:"payload"`
	TrackingId string `json:"trackingId"`
}

type RateFromResponse struct {
	Category     string `json:"category"`
	FromCurrency struct {
		Code    int    `json:"code"`
		Name    string `json:"name"`
		StrCode string `json:"strCode"`
	} `json:"fromCurrency"`
	ToCurrency struct {
		Code    int    `json:"code"`
		Name    string `json:"name"`
		StrCode string `json:"strCode"`
	} `json:"toCurrency"`
	Buy  float32 `json:"buy"`
	Sell float32 `json:"sell"`
}

func getCurrencyRates(queryParams map[string]string) (SuccessResponseFromTinkoffCurrencyRates, error) {
	req, err := http.NewRequest("GET", "https://api.tinkoff.ru/v1/currency_rates", nil)

	parsedResponse := SuccessResponseFromTinkoffCurrencyRates{}

	if err != nil {
		return parsedResponse, err
	}

	req.Header.Add("Content-Type", "application/json")
	query := req.URL.Query()

	for parameter, value := range queryParams {
		query.Add(parameter, value)
	}

	req.URL.RawQuery = query.Encode()

	httpClient := http.Client{}

	resp, err := httpClient.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return parsedResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return parsedResponse, err
	}

	err = json.Unmarshal(body, &parsedResponse)

	if err != nil {
		return parsedResponse, err
	}

	return parsedResponse, nil
}

func filterRates(rates []RateFromResponse, f func(rate RateFromResponse) bool) []RateFromResponse {
	filtered := make([]RateFromResponse, 0)

	for _, v := range rates {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
