package main

import (
	"github.com/jasonlvhit/gocron"
)

func registerCronJobs() {
	gocron.Every(1).Day().At("00:00:00").Do(updateBankRates)
	//gocron.Every(20).Seconds().Do(updateBankRates)

	<-gocron.Start()
}
