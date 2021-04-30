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
		categoryCondition, _ := inArray(rate.Category, []string{"C2CTransfers"})
		notZerosCondition := rate.Buy != 0 && rate.Sell != 0

		if notZerosCondition && categoryCondition {
			return true
		}

		return false
	})

	response.Payload.Rates = rates
	saveBankRates(convertTinkoffResponseToBankRateModels(response))
}

func convertTinkoffResponseToBankRateModels(resp SuccessResponseFromTinkoffCurrencyRates) []Rate {
	converted := make([]Rate, 0)

	for _, rate := range resp.Payload.Rates {
		bank, bankIsExist := findBankByAlias("tinkoff")
		fromCurrency, fromCurrencyIsExist := findCurrencyByAlias(rate.FromCurrency.Name)
		toCurrency, toCurrencyIsExist := findCurrencyByAlias(rate.ToCurrency.Name)

		if !bankIsExist || !fromCurrencyIsExist || !toCurrencyIsExist {
			continue
		}

		converted = append(converted, createBankRateModel(rate.Category, fromCurrency.Id, toCurrency.Id,
			resp.Payload.LastUpdate.Milliseconds/1000, bank.Id, rate.Buy, rate.Sell))
	}

	return converted
}
