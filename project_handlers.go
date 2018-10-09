package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project = &Project{}
	err := json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.createProject(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}
