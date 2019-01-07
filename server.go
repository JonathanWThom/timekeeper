package main

import (
	"database/sql"
	"github.com/gorilla/mux"
)

type server struct {
	db            *sql.DB
	router        *mux.Router
	currentUserID float64
}

func (s *server) init(dbName string) {
	db, err := sql.Open("postgres", "dbname="+dbName+" sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	s.db = db
	s.initKeys()
	s.router = mux.NewRouter()
	s.routes()
}
