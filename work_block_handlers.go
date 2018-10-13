package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// example workBlocksCreateHandler request
// curl -X POST \
//   http://localhost:8000/pay_periods/2/work_blocks \
//   -H 'Cache-Control: no-cache' \
//   -H 'Content-Type: application/json' \
//   -d '{ "project_id": 7, "started_at": "2006-01-02T15:04:05.000Z", "ended_at": "2006-01-02T15:04:05.000Z" }'

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

	err, hours := workBlock.hours()
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
