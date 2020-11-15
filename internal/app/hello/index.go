package hello

import "net/http"

func (s *Server) handleIndex() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Status string `json:"message"`
	}{
		Status: "Hello World",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}
