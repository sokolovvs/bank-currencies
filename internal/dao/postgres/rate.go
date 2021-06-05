package postgres

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	log "github.com/sirupsen/logrus"
	"github.com/sokolovvs/bank-currencies/internal/models"
	"github.com/sokolovvs/bank-currencies/pkg/database"
)

type RateDao struct {
}

func (*RateDao) Save(rate *models.Rate) error {
	jsonRate, _ := json.Marshal(rate)
	log.WithFields(log.Fields{"rate": string(jsonRate)}).Debug("Trying to Save Rate to database")

	if rate.Id == 0 {
		stmt, err := database.PgDb.Prepare("INSERT INTO rates (bank_id, category, from_currency_id, to_currency_id, buy, sell, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;")

		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")
			return err
		}

		defer stmt.Close()

		result := stmt.QueryRow(rate.BankId, rate.Category, rate.FromCurrencyId, rate.ToCurrencyId, rate.Buy, rate.Sell, rate.GetCreatedAtAsTime())

		if err = result.Scan(&rate.Id); err != nil {
			log.WithFields(log.Fields{"err": err, "rate": rate}).Warn(fmt.Sprintf("Error when inserting rate"))
		}

		return err
	}

	return errors.New("updating the Rate is not implemented")
}

func (*RateDao) SaveMany(rates []models.Rate) {
	rateDao := new(RateDao)

	for _, rate := range rates {
		err := rateDao.Save(&rate)

		if err != nil {
			log.Error(rate, " was not saved successfully")
		}
	}
}

type DtoFindRates struct {
	Limit int
}

func (*RateDao) FindByParams(dto DtoFindRates) ([]models.Rate, int, error) {
	rates := make([]models.Rate, 0)

	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "bank_id", "category", "from_currency_id", "to_currency_id", "buy", "sell", sb.As("CAST(EXTRACT(EPOCH FROM created_at) AS INTEGER)", "createdAtUTS"))
	sb.From("rates")

	sql, args := sb.Build()
	rows, err := database.PgDb.Query(sql, args...)

	if err != nil {
		log.Fatal(err)

		return rates, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		rate := models.Rate{}

		if err := rows.Scan(&rate.Id, &rate.BankId, &rate.Category, &rate.FromCurrencyId, &rate.ToCurrencyId, &rate.Buy, &rate.Sell, &rate.CreatedAt); err != nil {
			log.Fatal(err)

			return rates, 0, err
		}

		rates = append(rates, rate)
	}

	sb.Select(sb.As("COUNT(id)", "qty"))
	sqlCount, _ := sb.Build()
	var qty = 0
	result := database.PgDb.QueryRow(sqlCount)

	if err := result.Scan(&qty); err != nil {
		log.Fatal(err)

		return rates, qty, err
	}

	return rates, qty, err
}
