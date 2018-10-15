package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) usersCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user = &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.createUser(user)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	user.Token, err = user.getToken()
	user.Password = ""
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(user, w, r)
}
