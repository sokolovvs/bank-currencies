package main

func main() {
	var params map[string]string

	params = map[string]string{
		//"from": "RUB",
		//"to":   "USD",
	} // empty equal from/to for RUB, GPB, EUR, USD

	response, err := getCurrencyRates(params)

	if err != nil {
		panic(err)
	}

	rates := filterRates(response.Payload.Rates, func(rate RateFromResponse) bool {
		fromCondition, _ := inArray(rate.FromCurrency.Name, []string{"RUB", "EUR", "USD"})
		toCondition, _ := inArray(rate.ToCurrency.Name, []string{"RUB", "EUR", "USD"})

		if fromCondition && toCondition && rate.Category == "DebitCardsOperations" {
			return true
		}

		return false
	})

	response.Payload.Rates = rates
}
