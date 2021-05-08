package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	PgDb *sql.DB
)

/*
1st str - driver name

2nd str - connection data
*/
func getDatabaseSecrets() (string, string) {
	return os.Getenv("DB_DRIVER"), getConnectionStr()
}

func getConnectionStr() string {
	userName := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	host := os.Getenv("DB_HOST")

	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", userName, pass, dbName, sslMode, host)
}

func InitDb() {
	var err error

	PgDb, err = sql.Open(getDatabaseSecrets())

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Failed to create sql.DB instance")
		panic(err)
	}

	if err = PgDb.Ping(); err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Failed ping to database!")
		panic(err)
	}

	log.Info("Database ready to accept connections")
}
