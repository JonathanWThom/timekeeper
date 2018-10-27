package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createProject(project *Project) error {
	sql := `
		INSERT INTO projects(name, code, user_id)
		VALUES($1, $2, $3)
		RETURNING id, name, code, user_id
	`
	err := s.db.QueryRow(sql, project.Name, project.Code, project.UserID).
		Scan(&project.ID, &project.Name, &project.Code, &project.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) showProject(project *Project) error {
	sql := `
		SELECT id, name, code, user_id
		FROM projects
		WHERE id=$1
		AND user_id=$2
	`
	err := s.db.QueryRow(sql, project.ID, project.UserID).
		Scan(&project.ID, &project.Name, &project.Code, &project.UserID)
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
		AND user_id=$4
		RETURNING id, name, code, user_id
	`
	err := s.db.QueryRow(sql, project.Name, project.Code, project.ID, project.UserID).
		Scan(&project.ID, &project.Name, &project.Code, &project.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) deleteProject(project *Project) error {
	sql := `
		DELETE FROM projects
		WHERE id=$1
		AND user_id=$2
		RETURNING id, name, code, user_id
	`
	err := s.db.QueryRow(sql, project.ID, project.UserID).
		Scan(&project.ID, &project.Name, &project.Code, &project.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) indexProjects(userID int) ([]Project, error) {
	sql := `
		SELECT id, name, code, user_id
		FROM projects
		WHERE user_id=$1
	`
	rows, err := s.db.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var projects = []Project{}
	for rows.Next() {
		project := Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Code, &project.UserID)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
