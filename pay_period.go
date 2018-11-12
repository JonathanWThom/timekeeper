package main

import (
	"encoding/csv"
	"os"
	"time"
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

func (p *PayPeriod) generateReport(s *server) (string, error) {
	err := s.showPayPeriod(p)
	if err != nil {
		return "", err
	}

	records, err := p.buildCsv()
	if err != nil {
		return "", err
	}

	file, _ := os.OpenFile("report.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		return "", err
	}

	return "report.csv", nil
}

func (p *PayPeriod) buildCsv() ([][]string, error) {
	period := make(chan string, 1)
	dateRow := make(chan []string, 1)
	errs := make(chan error, 1)

	name := "Laura Syvertson" // eventually p.User.Name
	go p.getPeriod(period)
	go p.getDatesRow(dateRow, errs)

	err := <-errs
	if err != nil {
		return [][]string{}, err
	}

	records := [][]string{
		{"Name", name},
		{"Payroll Period", <-period},
		<-dateRow,
	}

	return records, nil
}

func (p *PayPeriod) getDatesRow(c chan<- []string, errs chan<- error) {
	layout := "2006-01-02T15:04:05Z"
	parsedStart, err := time.Parse(layout, p.StartedAt)
	if err != nil {
		errs <- err
		return
	}
	parsedEnd, err := time.Parse(layout, p.EndedAt)
	if err != nil {
		errs <- err
		return
	}
	dateToPrint := parsedStart
	dates := []string{"", "", "Date:", dateToPrint.Format("1/2")}

	for dateToPrint.Before(parsedEnd) {
		dateToPrint = dateToPrint.AddDate(0, 0, 1)
		dates = append(dates, dateToPrint.Format("1/2"))
	}
	dates = append(dates, "Totals")

	errs <- nil
	c <- dates
	close(c)
}

func (p *PayPeriod) getPeriod(c chan<- string) {
	c <- p.StartedAt[:10] + " - " + p.EndedAt[:10]
	close(c)
}
