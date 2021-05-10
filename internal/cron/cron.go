package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sokolovvs/bank-currencies/internal/service"
)

func RegisterCronJobs() {
	rateService := service.NewRateService()

	gocron.Every(1).Day().At("00:00:00").Do(rateService.UpdateBankRates)
	//gocron.Every(30).Seconds().Do(rateService.UpdateBankRates)

	<-gocron.Start()
}
