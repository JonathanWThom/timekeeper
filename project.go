package main

import "time"

// Project lives on the projects table
type Project struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Code   string `json:"code" db:"code"`
	UserID int    `json:"user_id" db:"user_id"`
}

func (p *Project) totalTimeOnDate(date time.Time) (float64, error) {

	return 0, nil
}
