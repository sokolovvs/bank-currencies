package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func registerServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloAction)
	r.HandleFunc("/api/v1/banks", getBanksAction).Methods("GET")
	r.HandleFunc("/api/v1/currencies", getCurrenciesAction).Methods("GET")

	appPort := os.Getenv("APP_PORT")
	err := http.ListenAndServe(fmt.Sprintf(":%s", appPort), r)

	if err != nil {
		log.Error(err)
	}
}
