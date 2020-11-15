package hello

import (
	"net/http"
)

func (s *Server) Authorize(authFunc func(r *http.Request) bool, handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if authFunc(r) {
			handle(w, r)
		} else {
			s.respond(w, r, struct{}{}, http.StatusUnauthorized)
		}
	}
}
