package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) payPeriodsCreateHandler(w http.ResponseWriter, r *http.Request) {
	var payPeriod = &PayPeriod{}
	// Working on making a proper request with dates here
	err := json.NewDecoder(r.Body).Decode(payPeriod)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.createPayPeriod(payPeriod)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(payPeriod, w, r)
}
