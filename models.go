package main

import "time"

// Project lives on the projects table
type Project struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Code string `json:"code" db:"code"`
}

// PayPeriod lives on the pay_periods table
type PayPeriod struct {
	ID        int    `json:"id" db:"id"`
	StartedOn string `json:"started_on" db:"started_on"` // stored as date in db
	EndedOn   string `json:"ended_on" db:"ended_on"`     // stored as date in db
	UserID    int    `json:"user_id" db:"user_id"`
}

// WorkBlock lives on the work_blocks table
type WorkBlock struct {
	ID          int    `json:"id" db:"id"`
	ProjectID   int    `json:"project_id" db:"project_id"`
	PayPeriodID int    `json:"pay_period_id" db:"pay_period_id"`
	Hours       int    `json:"hours" db:"hours"`
	StartedAt   string `json:"started_at" db:"started_at"` // store as timestamp in db - should I make it time here?
	EndedAt     string `json:"ended_at" db:"ended_at"`     // stored as timestamp in db
}

func (w *WorkBlock) hours() (int, error) {
	layout := "2006-01-02T15:04:05.000Z"
	end, err := time.Parse(layout, w.EndedAt)
	if err != nil {
		return 0, err
	}
	start, err := time.Parse(layout, w.StartedAt)
	if err != nil {
		return 0, err
	}

	return int(end.Sub(start)), nil
}
