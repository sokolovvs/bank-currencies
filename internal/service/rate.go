package service

import (
	"github.com/sokolovvs/bank-currencies/internal/service/external/exchange-rate/tinkoff"
)

type RateService struct {
	tinkoffExchangeRateUpdater *tinkoff.TinkoffExchangeRateUpdater
}

func NewRateService() *RateService {
	return &RateService{tinkoffExchangeRateUpdater: &tinkoff.TinkoffExchangeRateUpdater{}}
}

func (s *RateService) UpdateBankRates() {
	s.tinkoffExchangeRateUpdater.UpdateTinkoffRates()
}
