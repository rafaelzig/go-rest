package hello

import "net/http"

func (s *Server) initRoutes() {
	s.router.HandleFunc("/", s.Log(s.HandleIndex())).Methods(http.MethodGet)
	s.router.HandleFunc("/health", s.Log(s.HandleHealth())).Methods(http.MethodGet)
	s.router.HandleFunc("/shutdown", s.Log(s.Authorize(s.BasicAuth("admin", "password"), s.HandleShutdown()))).Methods(http.MethodDelete)
}
