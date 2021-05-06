package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func updateTinkoffRates() {
	log.Info(fmt.Sprintf("cron task %s was called", "updateTinkoffRates()"))

	defaultFilterFunc := func(rate RateFromResponse) bool {
		categoryCondition, _ := inArray(rate.Category, []string{"C2CTransfers"})
		notZerosCondition := rate.Buy != 0 && rate.Sell != 0

		if notZerosCondition && categoryCondition {
			return true
		}

		return false
	}

	updateTinkoffRatesByParams(map[string]string{"from": "USD", "to": "RUB"}, defaultFilterFunc)
	updateTinkoffRatesByParams(map[string]string{"from": "EUR", "to": "RUB"}, defaultFilterFunc)
	updateTinkoffRatesByParams(map[string]string{"from": "KZT", "to": "RUB"}, defaultFilterFunc)
	updateTinkoffRatesByParams(map[string]string{"from": "CAD", "to": "RUB"}, defaultFilterFunc)
	updateTinkoffRatesByParams(map[string]string{"from": "AUD", "to": "RUB"}, defaultFilterFunc)
}

func updateTinkoffRatesByParams(params map[string]string, filterFunc func(response RateFromResponse) bool) {
	response, err := getCurrencyRates(params)
	rateDao := new(RateDao)

	if err != nil {
		log.Error(err)
		return
	}

	rates := filterRates(response.Payload.Rates, filterFunc)

	response.Payload.Rates = rates
	rateDao.saveMany(convertTinkoffResponseToBankRateModels(response))
}

func convertTinkoffResponseToBankRateModels(resp SuccessResponseFromTinkoffCurrencyRates) []Rate {
	bankDao := new(BankDao)
	currencyDao := new(CurrencyDao)
	converted := make([]Rate, 0)

	for _, rate := range resp.Payload.Rates {
		bank, bankIsExist := bankDao.FindByAlias("tinkoff")
		fromCurrency, fromCurrencyIsExist := currencyDao.FindByAlias(rate.FromCurrency.Name)
		toCurrency, toCurrencyIsExist := currencyDao.FindByAlias(rate.ToCurrency.Name)

		if !bankIsExist || !fromCurrencyIsExist || !toCurrencyIsExist {
			continue
		}

		converted = append(converted, createBankRateModel(rate.Category, fromCurrency.Id, toCurrency.Id,
			resp.Payload.LastUpdate.Milliseconds/1000, bank.Id, rate.Buy, rate.Sell))
	}

	return converted
}
