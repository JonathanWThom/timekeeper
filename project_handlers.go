package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) projectsCreateHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *server) projectsShowHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *server) projectsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var project = &Project{}
	err := json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project.ID = id

	err = s.updateProject(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project := Project{ID: id}
	err = s.deleteProject(&project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectsIndexHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := s.indexProjects()
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(projects, w, r)
}
