package models

import "time"

type Rate struct {
	Id             int    `json:"id"`
	BankId         int    `json:"bank_id"`
	Category       string `json:"category"`
	FromCurrencyId int    `json:"from_currency_id"`
	ToCurrencyId   int    `json:"to_currency_id"`
	Buy            int    `json:"buy"`
	Sell           int    `json:"sell"`
	CreatedAt      int    `json:"created_at"` // unix timestamp
}

func CreateBankRateModel(category string, fromCurrencyId, toCurrencyId, createdAt, bankId, buy, sell int) Rate {
	return Rate{
		Category: category, FromCurrencyId: fromCurrencyId, ToCurrencyId: toCurrencyId,
		Buy: buy, Sell: sell, CreatedAt: createdAt, BankId: bankId,
	}
}

func (r *Rate) GetConvertedBuyRate() float32 {
	return float32(r.Buy / 100)
}

func (r *Rate) GetConvertedSellRate() float32 {
	return float32(r.Sell / 100)
}

func (r *Rate) GetCreatedAtAsTime() time.Time {
	return time.Unix(int64(r.CreatedAt), 0)
}
