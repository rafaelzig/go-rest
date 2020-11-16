package hello

import "net/http"

func (s *Server) initRoutes() {
	s.router.HandleFunc("/", s.Log(s.handleIndex())).Methods(http.MethodGet)
	s.router.HandleFunc("/health", s.Log(s.handleHealth())).Methods(http.MethodGet)
	s.router.HandleFunc("/shutdown", s.Log(s.authorize(s.basicAuth("admin", "password"), s.handleShutdown()))).Methods(http.MethodDelete)
}
