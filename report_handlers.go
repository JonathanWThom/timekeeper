package main

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// https://stackoverflow.com/questions/24116147/golang-how-to-download-file-in-browser-from-golang-server

func (s *server) reportsShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	payPeriodID, err := strconv.Atoi(vars["pay_period_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	payPeriod := &PayPeriod{ID: payPeriodID}
	userID, err := payPeriod.userID(s)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	payPeriod.UserID = int(userID)
	if s.currentUserID != userID {
		err = errors.New("Unauthorized")
		jsonUnauthorized(err, w, r)
		return
	}

	path, err := payPeriod.generateReport(s)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	http.ServeFile(w, r, path)
}
