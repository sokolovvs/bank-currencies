package postgres

import (
	log "github.com/sirupsen/logrus"
	"github.com/sokolovvs/bank-currencies/internal/models"
	"github.com/sokolovvs/bank-currencies/pkg/database"
)

type BankDao struct {
}

func (*BankDao) FindByAlias(alias string) (bank models.Bank, isExist bool) {
	stmt, err := database.PgDb.Prepare("SELECT id, alias FROM banks WHERE alias=$1;")

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")
		isExist = false
		return
	}

	defer stmt.Close()

	result := stmt.QueryRow(alias)

	if err := result.Scan(&bank.Id, &bank.Alias); err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error when trying to get Bank by alias " + alias)
		isExist = false
		return
	}

	isExist = true
	return
}

func (*BankDao) FindAll() ([]models.Bank, error) {
	banks := make([]models.Bank, 0)

	stmt, err := database.PgDb.Prepare("SELECT id, alias FROM banks")

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")

		return banks, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)

		return banks, err
	}

	defer rows.Close()

	for rows.Next() {
		bank := models.Bank{}

		if err := rows.Scan(&bank.Id, &bank.Alias); err != nil {
			log.Fatal(err)

			return banks, err
		}

		banks = append(banks, bank)
	}

	return banks, err
}
