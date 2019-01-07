package main

import (
	"encoding/csv"
	"fmt"
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

func (p *PayPeriod) projects(s *server) ([]Project, error) {
	sql := `
		SELECT projects.id, projects.name, projects.code, projects.user_id
		FROM projects
		INNER JOIN work_blocks
		ON work_blocks.project_id = projects.id
		INNER JOIN pay_periods
		ON pay_periods.id = $1
	`

	rows, err := s.db.Query(sql, p.ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var projects = []Project{}
	for rows.Next() {
		project := Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Code, &project.UserID)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func (p *PayPeriod) generateReport(s *server) (string, error) {
	err := s.showPayPeriod(p)
	if err != nil {
		return "", err
	}

	records, err := p.buildCsv(s)
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

func (p *PayPeriod) buildCsv(s *server) ([][]string, error) {
	period := make(chan string, 1)
	dateRow := make(chan []string, 1)
	projHeaderRow := make(chan []string, 1)
	projectRows := make(chan []string, 1)
	errs := make(chan error, 1)

	name := "Laura Syvertson" // eventually p.User.Name
	go p.getPeriod(period)
	go p.getDatesRow(dateRow, errs)
	go p.getProjHeaderRow(projHeaderRow, errs)
	go p.getProjectRows(s, projectRows, errs)

	err := <-errs
	if err != nil {
		return [][]string{}, err
	}

	records := [][]string{
		{"Name", name},
		{"Payroll Period", <-period},
		<-dateRow,
		<-projHeaderRow,
		<-projectRows,
	}

	return records, nil
}

func (p *PayPeriod) getDatesRow(c chan<- []string, errs chan<- error) {
	parsedStart, parsedEnd, err := p.getDates()
	if err != nil {
		errs <- err
		close(c)
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

func (p *PayPeriod) getDates() (time.Time, time.Time, error) {
	layout := "2006-01-02T15:04:05Z"
	parsedStart, err := time.Parse(layout, p.StartedAt)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	parsedEnd, err := time.Parse(layout, p.EndedAt)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return parsedStart, parsedEnd, nil
}

func (p *PayPeriod) getPeriod(c chan<- string) {
	c <- p.StartedAt[:10] + " - " + p.EndedAt[:10]
	close(c)
}

func (p *PayPeriod) getProjHeaderRow(c chan<- []string, errs chan<- error) {
	row := []string{"Proj #", "Project Name", "Service Item"}
	parsedStart, parsedEnd, err := p.getDates()
	if err != nil {
		errs <- err
		close(c)
		return
	}

	dateToPrint := parsedStart
	for dateToPrint.Before(parsedEnd) {
		dateToPrint = dateToPrint.AddDate(0, 0, 1)
		wkday := dateToPrint.Format("Mon")
		row = append(row, wkday)
	}

	c <- row
	close(c)
}

func (p *PayPeriod) dates() ([]time.Time, error) {
	start, end, err := p.getDates()
	if err != nil {
		return []time.Time{}, err
	}
	dateToPrint := start
	var dates []time.Time
	for dateToPrint.Before(end) {
		dateToPrint = dateToPrint.AddDate(0, 0, 1)
		dates = append(dates, dateToPrint)
	}

	return dates, nil
}

func (p *PayPeriod) getProjectRows(s *server, c chan<- []string, errs chan<- error) {
	projects, err := p.projects(s)

	if err != nil {
		errs <- err
		close(c)
		return
	}

	for _, project := range projects {
		// dates, err := p.dates()
		// if err != nil {
		// 	errs <- err
		// 	return
		// }
		// for _, date := range dates {
		// 	total, err := project.totalTimeOnDate(date)
		// 	if err != nil {
		// 		errs <- err
		// 		return
		// 	}
		// }
		fmt.Println(project)
	}
	// /// get total down here
	c <- []string{}
	close(c)
}
