package main

type BankRateModel struct {
	Id               int
	Category         string
	FromCurrencyName string
	ToCurrencyName   string
	Buy              float32
	Sell             float32
	LastUpdate       int
}
