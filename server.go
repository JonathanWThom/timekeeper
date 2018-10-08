package main

import "database/sql"

type server struct {
	db *sql.DB
}

func (s *server) init() {
	db, err := sql.Open("postgres", "dbname=timekeeper sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	s.db = db
}
