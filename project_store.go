package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createProject(project *Project) error {
	sql := `
		INSERT INTO projects(name, code)
		VALUES($1, $2)
		RETURNING id, name, code
	`
	err := s.db.QueryRow(sql, project.Name, project.Code).
		Scan(&project.ID, &project.Name, &project.Code)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) showProject(project *Project) error {
	sql := `
		SELECT id, name, code
		FROM projects
		WHERE id=$1
	`
	err := s.db.QueryRow(sql, project.ID).
		Scan(&project.ID, &project.Name, &project.Code)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) updateProject(project *Project) error {
	// Strip out nil values?

	sql := `
		UPDATE projects
		SET name=$1, code=$2
		WHERE id=$3
		RETURNING id, name, code
	`
	err := s.db.QueryRow(sql, project.Name, project.Code, project.ID).
		Scan(&project.ID, &project.Name, &project.Code)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) deleteProject(project *Project) error {
	sql := `
		DELETE FROM projects
		WHERE id=$1
		RETURNING id, name, code
	`
	err := s.db.QueryRow(sql, project.ID).
		Scan(&project.ID, &project.Name, &project.Code)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) indexProjects() ([]Project, error) {
	sql := `
		SELECT id, name, code
		FROM projects
	`
	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var projects = []Project{}
	for rows.Next() {
		project := Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Code)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
