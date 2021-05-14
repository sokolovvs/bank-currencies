package v1

import (
	"net/http"
	"testing"
)

func TestApiV1CurrenciesReturnsOkStatusCode(t *testing.T) {
	//TODO need to read domain name & port from env
	res, err := http.Get("http://localhost:8888/api/v1/currencies")

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Status code is not OK (200)")
	}
}
