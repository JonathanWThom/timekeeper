package main

func (s *server) routes() {
	// Projects
	s.router.HandleFunc("/projects", s.createProjectHandler).Methods("POST")
}
