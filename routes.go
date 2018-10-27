package main

func (s *server) routes() {
	// Projects
	s.router.HandleFunc(
		"/users/{user_id}/projects",
		s.validateTokenMiddleware(s.projectsCreateHandler)).
		Methods("POST")
	s.router.HandleFunc(
		"/users/{user_id}/projects/{id}",
		s.validateTokenMiddleware(s.projectsShowHandler)).
		Methods("GET")
	s.router.HandleFunc(
		"/users/{user_id}/projects/{id}",
		s.validateTokenMiddleware(s.projectsUpdateHandler)).
		Methods("PATCH")
	s.router.HandleFunc(
		"/users/{user_id}/projects/{id}",
		s.validateTokenMiddleware(s.projectsDeleteHandler)).
		Methods("DELETE")
	s.router.HandleFunc(
		"/users/{user_id}/projects",
		s.validateTokenMiddleware(s.projectsIndexHandler)).
		Methods("GET")

	// PayPeriods
	s.router.HandleFunc(
		"/users/{user_id}/pay_periods",
		s.validateTokenMiddleware(s.payPeriodsCreateHandler)).
		Methods("POST")
	s.router.HandleFunc(
		"/users/{user_id}/pay_periods/{id}",
		s.validateTokenMiddleware(s.payPeriodsShowHandler)).
		Methods("GET")
	s.router.HandleFunc(
		"/users/{user_id}/pay_periods/{id}",
		s.validateTokenMiddleware(s.payPeriodsUpdateHandler)).
		Methods("PATCH")
	s.router.HandleFunc(
		"/users/{user_id}/pay_periods/{id}",
		s.validateTokenMiddleware(s.payPeriodsDeleteHandler)).
		Methods("DELETE")
	s.router.HandleFunc(
		"/users/{user_id}/pay_periods",
		s.validateTokenMiddleware(s.payPeriodsIndexHandler)).
		Methods("GET")

	// WorkBlocks
	// project_id and user_id are both received from request body, not url params
	s.router.HandleFunc(
		"/pay_periods/{pay_period_id}/work_blocks",
		s.validateTokenMiddleware(s.workBlocksCreateHandler)).
		Methods("POST")
	s.router.HandleFunc(
		"/pay_periods/{pay_period_id}/work_blocks/{id}",
		s.validateTokenMiddleware(s.workBlocksShowHandler)).
		Methods("GET")
	s.router.HandleFunc(
		"/pay_periods/{pay_period_id}/work_blocks/{id}",
		s.validateTokenMiddleware(s.workBlocksUpdateHandler)).
		Methods("PATCH")
	s.router.HandleFunc(
		"/pay_periods/{pay_period_id}/work_blocks/{id}",
		s.validateTokenMiddleware(s.workBlocksDeleteHandler)).
		Methods("DELETE")
	s.router.HandleFunc(
		"/pay_periods/{pay_period_id}/work_blocks",
		s.validateTokenMiddleware(s.workBlocksIndexHandler)).
		Methods("GET")

	// Users
	s.router.HandleFunc(
		"/users",
		s.usersCreateHandler).
		Methods("POST")
	// 	s.router.HandleFunc("/users/{id}", s.usersUpdateHandler).Methods("PATCH")
	// 	s.router.HandleFunc("/users/{id}", s.usersDeleteHandler).Methods("DELETE")

	// Sessions
	s.router.HandleFunc(
		"/sessions",
		s.sessionsCreateHandler).
		Methods("POST")

		// Reports
	s.router.HandleFunc(
		"/pay_periods/{pay_period_id}/report",
		s.validateTokenMiddleware(s.reportsShowHandler)).
		Methods("GET")
}
