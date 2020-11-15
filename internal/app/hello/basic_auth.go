package hello

import "net/http"

func (s *Server) BasicAuth(a string, p string) func(*http.Request) bool {
	return func(r *http.Request) bool {
		auth, password, ok := r.BasicAuth()
		return ok && a == auth && p == password
	}
}
