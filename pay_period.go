package main

// PayPeriod lives on the pay_periods table
type PayPeriod struct {
	ID        int    `json:"id" db:"id"`
	StartedAt string `json:"started_at" db:"started_at"` // stored as date in db
	EndedAt   string `json:"ended_at" db:"ended_at"`     // stored as date in db
	UserID    int    `json:"user_id" db:"user_id"`
}
