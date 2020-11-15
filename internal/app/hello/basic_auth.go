package hello

import "net/http"

func (s *Server) BasicAuth(u string, p string) func(*http.Request) bool {
	return func(r *http.Request) bool {
		username, password, ok := r.BasicAuth()
		return ok && u == username && p == password
	}
}
