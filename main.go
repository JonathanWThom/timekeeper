package main

import (
	"github.com/jonathanwthom/timekeeper/report"
	"log"
	"net/http"
)

func main() {
	var s server
	s.init()
	report.Run()
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
