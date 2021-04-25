package main

import (
	"database/sql"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error loading .env file")
	}

	dbConnectionStr := os.Getenv("DB_CONNECTION")
	dbDriver := os.Getenv("DB_DRIVER")
	sql.Open(dbDriver, dbConnectionStr)

	gocron.Every(1).Day().At("00:00:00").Do(updateBankRates)

	<-gocron.Start()
}

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
