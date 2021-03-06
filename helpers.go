package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonSuccess(response interface{}, w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(response)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func jsonError(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Printf("Handling %q: %v", r.RequestURI, err)
}

func jsonUnauthorized(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), http.StatusUnauthorized)
	log.Printf("Handling %q: %v", r.RequestURI, err)
}
