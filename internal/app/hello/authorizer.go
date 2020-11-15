package hello

import (
	"net/http"
)

func (s *Server) checkAuthorization(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized(r) {
			h(w, r)
		} else {
			s.respond(w, r, struct{}{}, http.StatusUnauthorized)
		}
	}
}

func isAuthorized(r *http.Request) bool {
	auth, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return auth == "admin" && password == "password"
}
