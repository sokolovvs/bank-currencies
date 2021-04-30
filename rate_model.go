package main

type Rate struct {
	Id             int
	BankId         int
	Category       string
	FromCurrencyId int
	ToCurrencyId   int
	Buy            int
	Sell           int
	CreatedAt      int //unix timestamp
}

func createBankRateModel(category string, fromCurrencyId, toCurrencyId, createdAt, bankId int, buy, sell float32) Rate {
	return Rate{
		Category: category, FromCurrencyId: fromCurrencyId, ToCurrencyId: toCurrencyId,
		Buy: int(buy * 100), Sell: int(sell * 100), CreatedAt: createdAt, BankId: bankId,
	}
}
