package main

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
