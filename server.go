package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func registerServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		serializedResponse, _ := json.Marshal(map[string]string{"msg": "Welcome to this life-changing API.\nIts the best API, its true, all other API's are fake."})
		fmt.Fprintf(w, string(serializedResponse))
	})

	appPort := os.Getenv("APP_PORT")
	err := http.ListenAndServe(fmt.Sprintf(":%s", appPort), r)

	if err != nil {
		log.Error(err)
	}
}
