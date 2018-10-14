package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// curl -X POST \
//   http://localhost:8000/pay_periods/2/work_blocks \
//   -H 'Cache-Control: no-cache' \
//   -H 'Content-Type: application/json' \
//   -d '{ "project_id": 7, "started_at": "2018-02-10T15:04:05.000Z", "ended_at": "2018-03-08T15:04:05.000Z" }'

func (s *server) workBlocksCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	payPeriodID, err := strconv.Atoi(vars["pay_period_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	workBlock := WorkBlock{PayPeriodID: int(payPeriodID)}
	err = json.NewDecoder(r.Body).Decode(&workBlock)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	hours, err := workBlock.hours()
	if err != nil {
		jsonError(err, w, r)
		return
	}
	workBlock.Hours = hours

	err = s.createWorkBlock(&workBlock)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(workBlock, w, r)
}

func (s *server) workBlocksShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	payPeriodID, err := strconv.Atoi(vars["pay_period_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	workBlock := WorkBlock{PayPeriodID: int(payPeriodID), ID: int(id)}
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.showWorkBlock(&workBlock)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(workBlock, w, r)
}
