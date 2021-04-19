package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SuccessResponseFromTinkoffCurrencyRates struct {
	ResultCode string `json:"resultCode"`
	Payload    struct {
		LastUpdate struct {
			Milliseconds int `json:"milliseconds"`
		} `json:"lastUpdate"`
		Rates []struct {
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
		} `json:"rates"`
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
	} `json:"payload"`
	TrackingId string `json:"trackingId"`
}

func main() {
	req, err := http.NewRequest("GET", "https://api.tinkoff.ru/v1/currency_rates", nil)

	if err != nil {
		panic(err)
	}

	req.URL.Query().Add("from", "USD")
	req.URL.Query().Add("to", "RUB")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil || resp.StatusCode != 200 {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	parsedResponse := SuccessResponseFromTinkoffCurrencyRates{}

	err = json.Unmarshal(body, &parsedResponse)

	if err != nil {
		panic(err)
	}

	fmt.Println(parsedResponse)
}
