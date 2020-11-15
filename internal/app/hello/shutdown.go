package hello

import (
	"net/http"
	"syscall"
)

func (s *Server) HandleShutdown() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{
		Status: "shutdown initiated",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if s.errChan != nil {
			s.errChan <- syscall.SIGTERM
		}
		s.respond(w, r, response, http.StatusAccepted)
	}
}
