package main

import (
	"encoding/csv"
	"log"
	"os"
)

// PayPeriod lives on the pay_periods table
type PayPeriod struct {
	ID        int    `json:"id" db:"id"`
	StartedAt string `json:"started_at" db:"started_at"` // stored as date in db
	EndedAt   string `json:"ended_at" db:"ended_at"`     // stored as date in db
	UserID    int    `json:"user_id" db:"user_id"`
}

func (p *PayPeriod) userID(s *server) (float64, error) {
	sql := `
		SELECT user_id
		FROM pay_periods
		WHERE id=$1
	`
	err := s.db.QueryRow(sql, p.ID).Scan(&p.UserID)
	if err != nil {
		return 0, err
	}

	return float64(p.UserID), nil
}

func (p *PayPeriod) generateReport() (string, error) {

	// this date will be fill in by db values eventually
	records := [][]string{
		{"Name", "Laura Syvertson"},
		{"Payroll Period", "03-09-2018 - 03-23-2018"},
		{"", "", "Date:", "3/9", "3/10", "3/11", "3/12", "3/13", "3/14", "3/15", "3/16", "3/17", "3/18", "3/19", "3/20", "3/21", "3/22", "3/23", "Totals"},
	}

	// for _, block := range payPeriod.workPeriods {
	// 	record := []string{"eenie", "meenie", "minie", "mo"}
	// 	records = append(records, record)
	// }

	file, _ := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return "test.csv", nil
}
