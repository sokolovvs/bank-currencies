package main

import (
	log "github.com/sirupsen/logrus"
)

func updateTinkoffRates() {
	updateMainRatesTinkoff()
}

func updateMainRatesTinkoff() {
	var params map[string]string

	params = map[string]string{
		"from": "RUB",
		"to":   "USD",
	}

	response, err := getCurrencyRates(params)

	if err != nil {
		log.Error(err)
		return
	}

	rates := filterRates(response.Payload.Rates, func(rate RateFromResponse) bool {
		fromCondition, _ := inArray(rate.FromCurrency.Name, []string{"RUB", "EUR", "USD"})
		toCondition, _ := inArray(rate.ToCurrency.Name, []string{"RUB", "EUR", "USD"})
		//categoryCondition, _ := inArray(rate.Category, []string{"C2CTransfers"})
		notZerosCondition := rate.Buy != 0 && rate.Sell != 0

		if fromCondition && toCondition /*&& categoryCondition*/ && notZerosCondition {
			return true
		}

		return false
	})

	response.Payload.Rates = rates
	saveBankRates(convertTinkoffResponseToBankRateModels(response))
}

func convertTinkoffResponseToBankRateModels(resp SuccessResponseFromTinkoffCurrencyRates) []BankRateModel {
	converted := make([]BankRateModel, len(resp.Payload.Rates))

	for index, v := range resp.Payload.Rates {
		model := BankRateModel{
			Category: v.Category, FromCurrencyName: v.FromCurrency.Name, ToCurrencyName: v.ToCurrency.Name,
			Buy: v.Buy, Sell: v.Sell, LastUpdate: resp.Payload.LastUpdate.Milliseconds,
		}
		converted[index] = model
	}

	return converted
}
