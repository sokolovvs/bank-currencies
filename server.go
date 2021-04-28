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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to this life-changing API.\nIts the best API, its true, all other API's are fake.")
	})

	appPort := os.Getenv("APP_PORT")
	err := http.ListenAndServe(fmt.Sprintf(":%s", appPort), r)

	if err != nil {
		log.Error(err)
	}
}
