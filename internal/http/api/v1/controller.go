package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/sokolovvs/bank-currencies/internal/dao/postgres"
	"github.com/sokolovvs/bank-currencies/internal/models"
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

func (*HttpApiV1Controller) GetRates(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusBadRequest)

		serializedResponse, _ := json.Marshal(map[string]string{"message": "Bad request"})
		fmt.Fprintf(w, string(serializedResponse))

		return
	}

	dto := &postgres.DtoFindRates{}

	if err := schema.NewDecoder().Decode(dto, r.Form); err != nil {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusBadRequest)

		serializedResponse, _ := json.Marshal(map[string]string{"message": "Bad request"})
		fmt.Fprintf(w, string(serializedResponse))

		return
	}

	validate := validator.New()

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err := validate.Struct(dto)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		errorBody := make(map[string][]string)

		for _, err := range err.(validator.ValidationErrors) {
			if errs, exists := errorBody[err.StructField()]; exists {
				errorBody[err.StructField()] = append(errs, err.Error())
			} else {
				errs = make([]string, 0)
				errorBody[err.StructField()] = append(errs, err.Error())
			}
		}

		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		serializedResponse, _ := json.Marshal(map[string]interface{}{"message": "Request contains errors", "errors": errorBody})
		fmt.Fprintf(w, string(serializedResponse))

		return
	}

	rateDao := new(postgres.RateDao)
	rates, qty, err := rateDao.FindByParams(dto)
	serializedResponse, errJson := json.Marshal(struct {
		Qty   int           `json:"qty"`
		Rates []models.Rate `json:"rates"`
	}{
		Qty:   qty,
		Rates: rates,
	})

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
