package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func findCurrencyByAlias(alias string) (currency Currency, isExist bool) {
	stmt, err := db.Prepare("SELECT id, name, alias FROM currencies WHERE alias=$1;")

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")
		isExist = false
		return
	}

	defer stmt.Close()

	result := stmt.QueryRow(alias)

	if err := result.Scan(&currency.Id, &currency.Name, &currency.Alias); err != nil {
		log.WithFields(log.Fields{"err": err}).Warn(fmt.Sprintf("Error when trying to get Currency by alias " + alias))
		isExist = false
		return
	}

	isExist = true
	return
}
