package hello

import "net/http"

func (s *Server) HandleIndex() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hello World",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}
