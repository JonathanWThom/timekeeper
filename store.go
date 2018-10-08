package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createProject() (*Project, error) {
	// this might be passed in as a argument eventually
	project := Project{}

	sql := `
		INSERT INTO projects(name, code)
		VALUES($1, $2)
		RETURNING name, code, id
	`
	err := s.db.QueryRow(sql, "name", "code").
		Scan(&project.Name, &project.Code, &project.ID)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
