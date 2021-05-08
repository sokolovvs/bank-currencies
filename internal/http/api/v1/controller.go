package v1

import (
	"encoding/json"
	"fmt"
	"github.com/sokolovvs/bank-currencies/internal/dao/postgres"
	"net/http"
)

type HttpApiV1Controller struct {
}

func (*HttpApiV1Controller) GetBanks(w http.ResponseWriter, r *http.Request) {
	bankDao := new(postgres.BankDao)
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

func (*HttpApiV1Controller) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencyDao := new(postgres.CurrencyDao)
	banks, err := currencyDao.FindAll()
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
