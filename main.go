package main

import (
	"log"
	"net/http"
)

func main() {
	var server server
	server.init()
	log.Fatal(http.ListenAndServe(":8000", server.router))
}
