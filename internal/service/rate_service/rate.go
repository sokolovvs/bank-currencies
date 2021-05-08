package rate_service

import (
	"github.com/sokolovvs/bank-currencies/internal/service/external/exchange-rate/tinkoff"
)

func UpdateBankRates() {
	tinkoff.UpdateTinkoffRates()
}
