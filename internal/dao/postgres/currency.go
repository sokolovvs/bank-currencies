package postgres

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sokolovvs/bank-currencies/internal/models"
	"github.com/sokolovvs/bank-currencies/pkg/database"
)

type CurrencyDao struct {
}

func (*CurrencyDao) FindByAlias(alias string) (currency models.Currency, isExist bool) {
	stmt, err := database.PgDb.Prepare("SELECT id, name, alias FROM currencies WHERE alias=$1;")

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")
		isExist = false
		return
	}

	defer stmt.Close()

	result := stmt.QueryRow(alias)

	if err := result.Scan(&currency.Id, &currency.Name, &currency.Alias); err != nil {
		log.WithFields(log.Fields{"err": err}).Warn(fmt.Sprintf("Error when trying to get Currency by alias " + alias))
		isExist = false
		return
	}

	isExist = true
	return
}

func (*CurrencyDao) FindAll() ([]models.Currency, error) {
	currencies := make([]models.Currency, 0)

	stmt, err := database.PgDb.Prepare("SELECT id, name, alias FROM currencies")

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
		currency := models.Currency{}

		if err := rows.Scan(&currency.Id, &currency.Name, &currency.Alias); err != nil {
			log.Fatal(err)

			return currencies, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, err
}
