package main

import "time"

// WorkBlock lives on the work_blocks table
type WorkBlock struct {
	ID          int     `json:"id" db:"id"`
	ProjectID   int     `json:"project_id" db:"project_id"`
	PayPeriodID int     `json:"pay_period_id" db:"pay_period_id"`
	Hours       float64 `json:"hours" db:"hours"`
	StartedAt   string  `json:"started_at" db:"started_at"` // store as timestamp in db - should I make it time here?
	EndedAt     string  `json:"ended_at" db:"ended_at"`     // stored as timestamp in db
}

func (w *WorkBlock) userID(s *server) (float64, error) {
	sql := `
		SELECT user_id
		FROM pay_periods
		WHERE id=$1
	`
	var userID float64
	err := s.db.QueryRow(sql, w.PayPeriodID).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (w *WorkBlock) hours() (float64, error) {
	layout := "2006-01-02T15:04:05.000Z"
	end, err := time.Parse(layout, w.EndedAt)
	if err != nil {
		return 0, err
	}
	start, err := time.Parse(layout, w.StartedAt)
	if err != nil {
		return 0, err
	}

	return float64(end.Sub(start).Hours()), nil
}
