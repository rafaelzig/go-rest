package hello

import (
	"net/http"
	"syscall"
)

func (s *Server) handleShutdown() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.errChan <- syscall.SIGTERM
	}
}
