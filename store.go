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
	// strip out nil values?
	sql := `
		UPDATE projects
		SET column1=$1, column2=$2
		WHERE id=$3
		RETURNING id, name, code
	`
	err := s.db.QueryRow(sql, project.Name, project.Code).
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
	`
	_, err := s.db.Query(sql, project.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) indexProjects(projects []*Project) error {
	sql := `
		SELECT id, name, code
		FROM projects
	`
	rows, err := s.db.Query(sql)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		project := &Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Code)
		if err != nil {
			return err
		}

		projects = append(projects, project)
	}

	return nil
}
