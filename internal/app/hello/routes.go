package hello

import "net/http"

func (s *Server) initRoutes() {
	s.router.HandleFunc("/", s.logAccess(s.HandleIndex())).Methods(http.MethodGet)
	s.router.HandleFunc("/health", s.logAccess(s.HandleHealth())).Methods(http.MethodGet)
	s.router.HandleFunc("/shutdown", s.logAccess(s.checkAuthorization(s.HandleShutdown()))).Methods(http.MethodDelete)
}
