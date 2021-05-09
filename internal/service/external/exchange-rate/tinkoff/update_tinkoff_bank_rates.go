package tinkoff

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sokolovvs/bank-currencies/internal/dao/postgres"
	"github.com/sokolovvs/bank-currencies/internal/models"
	"github.com/sokolovvs/bank-currencies/pkg/utils"
)

func UpdateTinkoffRates() {
	log.Info(fmt.Sprintf("cron task %s was called", "updateTinkoffRates()"))

	defaultFilterFunc := func(rate RateFromResponse) bool {
		categoryCondition, _ := utils.InArray(rate.Category, []string{"C2CTransfers"})
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
	response, err := GetCurrencyRates(params)
	rateDao := new(postgres.RateDao)

	if err != nil {
		log.Error("Fetching currencies was failed: ", err)
		return
	}

	rates := FilterRates(response.Payload.Rates, filterFunc)

	response.Payload.Rates = rates
	rateDao.SaveMany(convertTinkoffResponseToBankRateModels(response))
}

func convertTinkoffResponseToBankRateModels(resp SuccessResponseFromTinkoffCurrencyRates) []models.Rate {
	bankDao := new(postgres.BankDao)
	currencyDao := new(postgres.CurrencyDao)
	converted := make([]models.Rate, 0)

	for _, r := range resp.Payload.Rates {
		bank, bankIsExist := bankDao.FindByAlias("tinkoff")
		fromCurrency, fromCurrencyIsExist := currencyDao.FindByAlias(r.FromCurrency.Name)
		toCurrency, toCurrencyIsExist := currencyDao.FindByAlias(r.ToCurrency.Name)

		if !bankIsExist || !fromCurrencyIsExist || !toCurrencyIsExist {
			continue
		}

		converted = append(converted, models.CreateBankRateModel(r.Category, fromCurrency.Id, toCurrency.Id,
			resp.Payload.LastUpdate.Milliseconds/1000, bank.Id, r.Buy, r.Sell))
	}

	return converted
}
