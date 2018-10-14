package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createWorkBlock(workBlock *WorkBlock) error {
	sql := `
		INSERT INTO work_blocks(project_id, pay_period_id, hours, started_at, ended_at)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, project_id, pay_period_id, hours, started_at, ended_at
	`
	err := s.db.QueryRow(
		sql,
		workBlock.ProjectID,
		workBlock.PayPeriodID,
		workBlock.Hours,
		workBlock.StartedAt,
		workBlock.EndedAt).
		Scan(
			&workBlock.ID,
			&workBlock.ProjectID,
			&workBlock.PayPeriodID,
			&workBlock.Hours,
			&workBlock.StartedAt,
			&workBlock.EndedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) showWorkBlock(workBlock *WorkBlock) error {
	sql := `
		SELECT *
		FROM work_blocks
		WHERE id=$1
		AND pay_period_id=$2
	`
	err := s.db.QueryRow(sql, workBlock.ID, workBlock.PayPeriodID).
		Scan(
			&workBlock.ID,
			&workBlock.ProjectID,
			&workBlock.PayPeriodID,
			&workBlock.Hours,
			&workBlock.StartedAt,
			&workBlock.EndedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) updateWorkBlock(workBlock *WorkBlock) error {
	sql := `
		UPDATE work_blocks
		SET project_id=$1, hours=$2, started_at=$3, ended_at=$4
		WHERE id=$5
		AND pay_period_id=$6
		RETURNING id, project_id, pay_period_id, hours, started_at, ended_at
	`
	err := s.db.QueryRow(
		sql,
		workBlock.ProjectID,
		workBlock.Hours,
		workBlock.StartedAt,
		workBlock.EndedAt,
		workBlock.ID,
		workBlock.PayPeriodID).
		Scan(
			&workBlock.ID,
			&workBlock.ProjectID,
			&workBlock.PayPeriodID,
			&workBlock.Hours,
			&workBlock.StartedAt,
			&workBlock.EndedAt)
	if err != nil {
		return err
	}

	return nil
}
