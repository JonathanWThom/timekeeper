package main

import (
	"log"
	"net/http"
)

func main() {
	var s server
	s.init()
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
