package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sokolovvs/bank-currencies/internal/service/rate_service"
)

func RegisterCronJobs() {
	gocron.Every(1).Day().At("00:00:00").Do(rate_service.UpdateBankRates)
	//gocron.Every(30).Seconds().Do(rate_service.UpdateBankRates)

	<-gocron.Start()
}
