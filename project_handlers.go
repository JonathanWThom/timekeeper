package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) projectsCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project := &Project{UserID: userID}
	err = json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	if s.currentUserID != float64(userID) {
		err = errors.New("Unauthorized")
		jsonUnauthorized(err, w, r)
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
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project := Project{UserID: userID, ID: int(id)}

	if s.currentUserID != float64(userID) {
		err = errors.New("Unauthorized")
		jsonUnauthorized(err, w, r)
		return
	}

	err = s.showProject(&project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project := &Project{UserID: userID, ID: id}
	err = json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	if s.currentUserID != float64(userID) {
		err = errors.New("Unauthorized")
		jsonUnauthorized(err, w, r)
		return
	}

	err = s.updateProject(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(project, w, r)
}

func (s *server) projectsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	project := &Project{UserID: userID, ID: id}
	err = json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	if s.currentUserID != float64(userID) {
		err = errors.New("Unauthorized")
		jsonUnauthorized(err, w, r)
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
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		jsonError(err, w, r)
		return
	}

	if s.currentUserID != float64(userID) {
		err = errors.New("Unauthorized")
		jsonUnauthorized(err, w, r)
		return
	}

	projects, err := s.indexProjects(userID)
	if err != nil {
		jsonError(err, w, r)
		return
	}

	jsonSuccess(projects, w, r)
}
