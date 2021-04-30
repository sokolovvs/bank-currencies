package main

import (
	log "github.com/sirupsen/logrus"
)

func findBankByAlias(alias string) (bank Bank, isExist bool) {
	stmt, err := db.Prepare("SELECT id, alias FROM banks WHERE alias=$1;")

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error when trying to prepare statement")
		isExist = false
		return
	}

	defer stmt.Close()

	result := stmt.QueryRow(alias)

	if err := result.Scan(&bank.Id, &bank.Alias); err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error when trying to get Bank by alias " + alias)
		isExist = false
		return
	}

	isExist = true
	return
}
