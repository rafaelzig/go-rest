package hello

import "net/http"

func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.Log(
		s.handleIndex())).Methods(http.MethodGet)
	s.Router.HandleFunc("/health", s.Log(
		s.handleHealth())).Methods(http.MethodGet)
	s.Router.HandleFunc("/shutdown", s.Log(s.authorize(s.basicAuth("admin", "password"),
		s.handleShutdown()))).Methods(http.MethodDelete)
}
