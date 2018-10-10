package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) projectCreateHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *server) projectShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project := Project{ID: int(id)}
	err = s.showProject(&project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var project = &Project{}
	err := json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.updateProject(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var project = &Project{}
	err := json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	err = s.deleteProject(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectsIndexHandler(w http.ResponseWriter, r *http.Request) {
	var projects = []*Project{}
	err := s.indexProjects(projects)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(projects, w, r)
}
