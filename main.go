package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	bootstrap()
	go registerCronJobs()
	registerServer()
}

func bootstrap() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error loading .env file")
	}
}
