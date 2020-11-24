package hello

import "net/http"

func (s *Server) Routes() {
	s.Router.Use(s.LoggingMiddleware())
	s.Router.HandleFunc("/", s.handleIndex()).Methods(http.MethodGet)
	s.Router.HandleFunc("/health", s.handleHealth()).Methods(http.MethodGet)
	s.Router.HandleFunc("/admin", s.WithAuth(s.basicAuth("admin", "password"),
		s.handleAdmin())).Methods(http.MethodGet)
}
