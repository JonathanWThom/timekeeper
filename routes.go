package main

func (s *server) routes() {
	// Projects
	s.router.HandleFunc("/projects", s.projectsCreateHandler).Methods("POST")
	s.router.HandleFunc("/projects/{id}", s.projectsShowHandler).Methods("GET")
	s.router.HandleFunc("/projects/{id}", s.projectsUpdateHandler).Methods("PATCH")
	s.router.HandleFunc("/projects/{id}", s.projectsDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/projects", s.projectsIndexHandler).Methods("GET")

	// PayPeriods
	s.router.HandleFunc("/users/{user_id}/pay_periods", s.payPeriodsCreateHandler).Methods("POST")
	s.router.HandleFunc("/users/{user_id}/pay_periods/{id}", s.payPeriodsShowHandler).Methods("GET")
	s.router.HandleFunc("/users/{user_id}/pay_periods/{id}", s.payPeriodsUpdateHandler).Methods("PATCH")
	s.router.HandleFunc("/users/{user_id}/pay_periods/{id}", s.payPeriodsDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/users/{user_id}/pay_periods", s.payPeriodsIndexHandler).Methods("GET")

	// WorkBlocks
	// project_id and user_id are both received from request body, not url params
	s.router.HandleFunc("/pay_periods/{pay_period_id}/work_blocks", s.workBlocksCreateHandler).Methods("POST")
	s.router.HandleFunc("/pay_periods/{pay_period_id}/work_blocks/{id}", s.workBlocksShowHandler).Methods("GET")
	s.router.HandleFunc("/pay_periods/{pay_period_id}/work_blocks/{id}", s.workBlocksUpdateHandler).Methods("PATCH")
	s.router.HandleFunc("/pay_periods/{pay_period_id}/work_blocks/{id}", s.workBlocksDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/pay_periods/{pay_period_id}/work_blocks", s.workBlocksIndexHandler).Methods("GET")

	// Users
	s.router.HandleFunc("/users", s.usersCreateHandler).Methods("POST")
	// 	s.router.HandleFunc("/users/{id}", s.usersUpdateHandler).Methods("PATCH")
	// 	s.router.HandleFunc("/users/{id}", s.usersDeleteHandler).Methods("DELETE")

	// Sessions
	s.router.HandleFunc("/sessions", s.sessionsCreateHandler).Methods("POST")
	// s.router.HandleFunc("/sessions", s.sessionsDeleteHandler).Methods("DELETE")
}
