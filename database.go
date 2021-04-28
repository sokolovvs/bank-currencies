package main

import (
	"fmt"
	"os"
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
	dbName := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSL_MODE")

	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", userName, pass, dbName, sslMode)
}
