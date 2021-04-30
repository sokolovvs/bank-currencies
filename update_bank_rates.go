package main

func updateBankRates() {
	updateTinkoffRates()
}

func saveBankRates(rates []Rate) {
	for _, rate := range rates {
		saveRate(rate)
	}
}
