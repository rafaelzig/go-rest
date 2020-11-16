package hello

import (
	"net/http"
)

func (s *Server) Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.info().Printf("%s request to %s\n", r.RequestURI, r.Method)
		h(w, r)
	}
}
