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
	ID        int       `json:"id" db:"id"`
	StartedOn time.Time `json:"started_on" db:"started_on"`
	EndedOn   time.Time `json:"ended_on" db:"ended_on"`
	UserID    int       `json:"user_id" db:"user_id"`
}
