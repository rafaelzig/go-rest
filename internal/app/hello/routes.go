package hello

import "net/http"

func (s *Server) initRoutes() {
	s.router.HandleFunc("/", s.logAccess(s.handleIndex())).Methods(http.MethodGet)
	s.router.HandleFunc("/health", s.logAccess(s.handleHealth())).Methods(http.MethodGet)
	s.router.HandleFunc("/shutdown", s.logAccess(s.checkAuthorization(s.handleShutdown()))).Methods(http.MethodDelete)
}
