package main

import (
	log "github.com/sirupsen/logrus"
)

func updateBankRates() {
	updateTinkoffRates()
}

func updateTinkoffRates() {
	updateMainRatesTinkoff()
}

func updateMainRatesTinkoff() {
	var params map[string]string

	params = map[string]string{
		//"from": "RUB",
		//"to":   "USD",
	} // empty equal from/to for RUB, GPB, EUR, USD

	response, err := getCurrencyRates(params)

	if err != nil {
		log.Error(err)
		return
	}

	rates := filterRates(response.Payload.Rates, func(rate RateFromResponse) bool {
		fromCondition, _ := inArray(rate.FromCurrency.Name, []string{"RUB", "EUR", "USD"})
		toCondition, _ := inArray(rate.ToCurrency.Name, []string{"RUB", "EUR", "USD"})

		if fromCondition && toCondition && rate.Category == "DebitCardsOperations" {
			return true
		}

		return false
	})

	response.Payload.Rates = rates
}
