package main

type Rate struct {
	Id             int    `json:"id"`
	BankId         int    `json:"bank_id"`
	Category       string `json:"category"`
	FromCurrencyId int    `json:"from_currency_id"`
	ToCurrencyId   int    `json:"to_currency_id"`
	Buy            int    `json:"buy"`
	Sell           int    `json:"sell"`
	CreatedAt      int    `json:"created_at"`
}

func createBankRateModel(category string, fromCurrencyId, toCurrencyId, createdAt, bankId int, buy, sell float32) Rate {
	return Rate{
		Category: category, FromCurrencyId: fromCurrencyId, ToCurrencyId: toCurrencyId,
		Buy: int(buy * 100), Sell: int(sell * 100), CreatedAt: createdAt, BankId: bankId,
	}
}
