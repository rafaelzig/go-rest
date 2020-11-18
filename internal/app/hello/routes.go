package hello

import "net/http"

func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.WithLog(
		s.handleIndex())).Methods(http.MethodGet)
	s.Router.HandleFunc("/health", s.WithLog(
		s.handleHealth())).Methods(http.MethodGet)
	s.Router.HandleFunc("/admin", s.WithLog(s.WithAuth(s.basicAuth("admin", "password"),
		s.handleAdmin()))).Methods(http.MethodGet)
}
