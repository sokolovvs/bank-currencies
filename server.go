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

	apiV1 := new(HttpApiV1Controller)

	r.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "running")
	})
	r.HandleFunc("/api/v1/banks", apiV1.GetBanks).Methods("GET")
	r.HandleFunc("/api/v1/currencies", apiV1.GetCurrencies).Methods("GET")

	appPort := os.Getenv("APP_PORT")
	err := http.ListenAndServe(fmt.Sprintf(":%s", appPort), r)

	if err != nil {
		log.Error(err)
	}
}
