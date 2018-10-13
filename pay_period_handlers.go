package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) payPeriodsCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	payPeriod := PayPeriod{UserID: int(userID)}
	err = json.NewDecoder(r.Body).Decode(&payPeriod)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.createPayPeriod(&payPeriod)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(payPeriod, w, r)
}
