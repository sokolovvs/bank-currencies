package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func saveRate(rate Rate) error {
	jsonRate, _ := json.Marshal(rate)
	log.WithFields(log.Fields{"rate": string(jsonRate)}).Debug("Trying to save Rate to database")

	if rate.Id == 0 {
		stmt, err := db.Prepare("INSERT INTO rates (bank_id, category, from_currency_id, to_currency_id, buy, sell, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;")

		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")
			return err
		}

		defer stmt.Close()

		result := stmt.QueryRow(rate.BankId, rate.Category, rate.FromCurrencyId, rate.ToCurrencyId, rate.Buy, rate.Sell, getCreatedAtAsTime(rate))

		if err = result.Scan(&rate.Id); err != nil {
			log.WithFields(log.Fields{"err": err, "rate": rate}).Warn(fmt.Sprintf("Error when inserting rate"))
		}

		return err
	}

	return errors.New("updating the Rate is not supported")
}

func findCurrencies() ([]Currency, error) {
	currencies := make([]Currency, 0)

	stmt, err := db.Prepare("SELECT id, name, alias FROM currencies")

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")

		return currencies, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)

		return currencies, err
	}

	defer rows.Close()

	for rows.Next() {
		currency := Currency{}

		if err := rows.Scan(&currency.Id, &currency.Name, &currency.Alias); err != nil {
			log.Fatal(err)

			return currencies, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, err
}
