package main

func (s *server) routes() {
	// Projects
	s.router.HandleFunc("/projects", s.projectCreateHandler).Methods("POST")
	s.router.HandleFunc("/projects/{id}", s.projectShowHandler).Methods("GET")
	s.router.HandleFunc("/projects/{id}", s.projectUpdateHandler).Methods("PATCH")
}
