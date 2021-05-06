package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	serializedResponse, _ := json.Marshal(map[string]string{"msg": "Welcome to this life-changing API.\nIts the best API, its true, all other API's are fake."})
	fmt.Fprintf(w, string(serializedResponse))
}

func getBanksAction(w http.ResponseWriter, r *http.Request) {
	bankDao := new(BankDao)
	banks, err := bankDao.FindAll()
	serializedResponse, errJson := json.Marshal(banks)

	if err != nil || errJson != nil {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusInternalServerError)

		serializedResponse, _ = json.Marshal(map[string]string{"message": "Internal server error"})
		fmt.Fprintf(w, string(serializedResponse))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(serializedResponse))
}

func getCurrenciesAction(w http.ResponseWriter, r *http.Request) {
	banks, err := findCurrencies()
	serializedResponse, errJson := json.Marshal(banks)

	if err != nil || errJson != nil {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusInternalServerError)

		serializedResponse, _ = json.Marshal(map[string]string{"message": "Internal server error"})
		fmt.Fprintf(w, string(serializedResponse))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(serializedResponse))
}
