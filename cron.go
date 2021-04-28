package main

import (
	"github.com/jasonlvhit/gocron"
)

func registerCronJobs() {
	gocron.Every(1).Day().Do(updateBankRates)

	<-gocron.Start()
}
