package main

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

func createBankRateModel(category string, fromCurrencyId, toCurrencyId, createdAt, bankId int, buy, sell float32) Rate {
	return Rate{
		Category: category, FromCurrencyId: fromCurrencyId, ToCurrencyId: toCurrencyId,
		Buy: int(buy * 100), Sell: int(sell * 100), CreatedAt: createdAt, BankId: bankId,
	}
}

func getConvertedBuyRate(r Rate) float32 {
	return float32(r.Buy / 100)
}

func getConvertedSellRate(r Rate) float32 {
	return float32(r.Sell / 100)
}

func getCreatedAtAsTime(r Rate) time.Time {
	return time.Unix(int64(r.CreatedAt), 0)
}
