package main

import (
	"reflect"
	"testing"
)

func TestBuildCsv(t *testing.T) {
	p := &PayPeriod{
		StartedAt: "2018-03-09T00:00:00Z",
		EndedAt:   "2018-03-16T00:00:00Z",
		ID:        3,
	}
	// need to stub out the payPeriod's workBlocks and db connection

	tests := []struct {
		payPeriod     *PayPeriod
		nameRow       []string
		periodRow     []string
		dateRow       []string
		projHeaderRow []string
	}{
		{
			p,
			[]string{"Name", "Laura Syvertson"},
			[]string{"Payroll Period", "2018-03-09 - 2018-03-16"},
			[]string{"", "", "Date:", "3/9", "3/10", "3/11", "3/12", "3/13", "3/14", "3/15", "3/16", "Totals"},
			[]string{"Proj #", "Project Name", "Service Item", "Sat", "Sun", "Mon", "Tues", "Wed", "Thu", "Fri"},
		},
	}

	var s server
	s.init("timekeeper_test")

	for _, test := range tests {
		records, _ := test.payPeriod.buildCsv(&s)
		nameRow, periodRow, dateRow, projHeaderRow := records[0], records[1], records[2], records[3]

		if !reflect.DeepEqual(nameRow, test.nameRow) {
			t.Errorf("Name Row was incorrect, expected %v, got %v", nameRow, test.nameRow)
		}

		if !reflect.DeepEqual(periodRow, test.periodRow) {
			t.Errorf("Period Row was incorrect, expected %v, got %v", periodRow, test.periodRow)
		}

		if !reflect.DeepEqual(dateRow, test.dateRow) {
			t.Errorf("Date Row was incorrect, expected %v, got %v", dateRow, test.dateRow)
		}

		if !reflect.DeepEqual(dateRow, test.dateRow) {
			t.Errorf("Project Header Row was incorrect, expected %v, got %v",
				projHeaderRow, test.projHeaderRow)
		}
	}
}

func BenchmarkBuildCsv(b *testing.B) {
	p := &PayPeriod{
		StartedAt: "2018-03-09T00:00:00Z",
		EndedAt:   "2018-03-16T00:00:00Z",
	}

	var s server
	s.init("timekeeper_test")
	for i := 0; i < b.N; i++ {
		p.buildCsv(&s)
	}
}

func setUpTearDownDb() {
	// TODO
}
