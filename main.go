package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB
)

func main() {
	bootstrap()
	go registerCronJobs()
	registerServer()
}

func bootstrap() {
	log.SetFormatter(&log.JSONFormatter{})
	loadEnvParams()
	initDb()
}

func initDb() {
	var err error

	db, err = sql.Open(getDatabaseSecrets())

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Failed to create sql.DB instance")
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Failed ping to database!")
		panic(err)
	}

	log.Info("Database ready to accept connections")
}

func loadEnvParams() {
	err := godotenv.Load()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error loading .env file")
	}
}
