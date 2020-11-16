package hello

import "net/http"

func (s *Server) initRoutes() {
	s.router.HandleFunc("/", s.Log(s.HandleIndex())).Methods(http.MethodGet)
	s.router.HandleFunc("/health", s.Log(s.handleHealth())).Methods(http.MethodGet)
	s.router.HandleFunc("/shutdown", s.Log(s.authorize(s.basicAuth("admin", "password"), s.HandleShutdown()))).Methods(http.MethodDelete)
}
