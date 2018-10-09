package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createProject(project *Project) error {
	sql := `
		INSERT INTO projects(name, code)
		VALUES($1, $2)
		RETURNING name, code, id
	`
	err := s.db.QueryRow(sql, project.Name, project.Code).
		Scan(&project.Name, &project.Code, &project.ID)
	if err != nil {
		return err
	}

	return nil
}
