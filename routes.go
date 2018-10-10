package main

func (s *server) routes() {
	// Projects
	s.router.HandleFunc("/projects", s.projectsCreateHandler).Methods("POST")
	s.router.HandleFunc("/projects/{id}", s.projectsShowHandler).Methods("GET")
	s.router.HandleFunc("/projects/{id}", s.projectsUpdateHandler).Methods("PATCH")
	s.router.HandleFunc("/projects/{id}", s.projectsDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/projects", s.projectsIndexHandler).Methods("GET")

	// PayPeriods
	s.router.HandleFunc("/pay_periods", s.payPeriodsCreateHandler).Methods("POST")
}
