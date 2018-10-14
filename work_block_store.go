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
