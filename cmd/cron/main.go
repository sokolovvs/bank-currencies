package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/sokolovvs/bank-currencies/internal/cron"
	"github.com/sokolovvs/bank-currencies/pkg/database"
)

func main() {
	bootstrap()
	cron.RegisterCronJobs()
}

func bootstrap() {
	log.SetFormatter(&log.JSONFormatter{})
	loadEnvParams()
	database.InitDb()
}

func loadEnvParams() {
	err := godotenv.Load()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error loading .env file")
	}
}
