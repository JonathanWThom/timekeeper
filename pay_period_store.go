package main

import (
	_ "github.com/lib/pq"
)

func (s *server) createPayPeriod(payPeriod *PayPeriod) error {
	sql := `
		INSERT INTO pay_periods(started_at, ended_at, user_id)
		VALUES($1, $2, $3)
		RETURNING id, started_at, ended_at, user_id
	`
	err := s.db.QueryRow(sql, payPeriod.StartedAt, payPeriod.EndedAt, payPeriod.UserID).
		Scan(&payPeriod.ID, &payPeriod.StartedAt, &payPeriod.EndedAt, &payPeriod.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) showPayPeriod(payPeriod *PayPeriod) error {
	sql := `
		SELECT id, started_at, ended_at, user_id
		FROM pay_periods
		WHERE user_id=$1
		AND id=$2
	`
	err := s.db.QueryRow(sql, payPeriod.UserID, payPeriod.ID).
		Scan(&payPeriod.ID, &payPeriod.StartedAt, &payPeriod.EndedAt, &payPeriod.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) updatePayPeriod(payPeriod *PayPeriod) error {
	// Strip out nil values?
	sql := `
		UPDATE pay_periods
		SET started_at=$1, ended_at=$2
		WHERE user_id=$3
		AND id=$4
		RETURNING id, started_at, ended_at, user_id
	`
	err := s.db.QueryRow(sql, payPeriod.StartedAt, payPeriod.EndedAt, payPeriod.UserID, payPeriod.ID).
		Scan(&payPeriod.ID, &payPeriod.StartedAt, &payPeriod.EndedAt, &payPeriod.UserID)
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
		RETURNING id, started_at, ended_at, user_id
	`
	err := s.db.QueryRow(sql, payPeriod.UserID, payPeriod.ID).
		Scan(&payPeriod.ID, &payPeriod.StartedAt, &payPeriod.EndedAt, &payPeriod.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) indexPayPeriods(userID int) ([]PayPeriod, error) {
	sql := `
		SELECT id, started_at, ended_at, user_id
		FROM pay_periods
		WHERE user_id=$1
	`
	rows, err := s.db.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var payPeriods = []PayPeriod{}
	for rows.Next() {
		payPeriod := PayPeriod{}
		err := rows.Scan(&payPeriod.ID, &payPeriod.StartedAt, &payPeriod.EndedAt, &payPeriod.UserID)
		if err != nil {
			return nil, err
		}

		payPeriods = append(payPeriods, payPeriod)
	}

	return payPeriods, nil
}
