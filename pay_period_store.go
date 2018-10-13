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

func (s *server) showPayPeriod(payPeriod *PayPeriod) error {
	sql := `
		SELECT id, started_on, ended_on, user_id
		FROM pay_periods
		WHERE user_id=$1
		AND id=$2
	`
	err := s.db.QueryRow(sql, payPeriod.UserID, payPeriod.ID).
		Scan(&payPeriod.ID, &payPeriod.StartedOn, &payPeriod.EndedOn, &payPeriod.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) updatePayPeriod(payPeriod *PayPeriod) error {
	// Strip out nil values?

	sql := `
		UPDATE pay_periods
		SET started_on=$1, ended_on=$2
		WHERE user_id=$3
		AND id=$4
		RETURNING id, started_on, ended_on, user_id
	`
	err := s.db.QueryRow(sql, payPeriod.StartedOn, payPeriod.EndedOn, payPeriod.UserID, payPeriod.ID).
		Scan(&payPeriod.ID, &payPeriod.StartedOn, &payPeriod.EndedOn, &payPeriod.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) deletePayPeriod(payPeriod *PayPeriod) error {
	sql := `
		DELETE FROM pay_periods
		WHERE user_id=$1
		AND id=$2
	`
	_, err := s.db.Query(sql, payPeriod.UserID, payPeriod.ID)
	if err != nil {
		return err
	}

	return nil
}
