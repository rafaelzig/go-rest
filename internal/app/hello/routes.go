package hello

func (s *Server) initRoutes() {
	s.router.HandleFunc("/", s.logAccess(s.handleIndex()))
	s.router.HandleFunc("/health", s.logAccess(s.handleHealth()))
	s.router.HandleFunc("/shutdown", s.logAccess(s.checkAuthorization(s.handleShutdown())))
}
