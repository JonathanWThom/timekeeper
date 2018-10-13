package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createPayPeriod(payPeriod *PayPeriod) error {
	sql := `
		INSERT INTO pay_periods(started_on, ended_on, user_id)
		VALUES($1, $2, $3)
		RETURNING id, started_on, ended_on, user_id
	`
	err := s.db.QueryRow(sql, payPeriod.StartedOn, payPeriod.EndedOn, payPeriod.UserID).
		Scan(&payPeriod.ID, &payPeriod.StartedOn, &payPeriod.EndedOn, &payPeriod.UserID)
	if err != nil {
		return err
	}

	return nil
}
