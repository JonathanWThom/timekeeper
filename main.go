package main

import "fmt"

func main() {
	var server server
	server.init()

	// TODO: REMOVE THIS
	project, err := server.createProject()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(project)
	}
}
