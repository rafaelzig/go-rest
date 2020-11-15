package hello

import (
	"log"
	"net/http"
)

func (s *Server) Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request to %s\n", r.RequestURI, r.Method)
		h(w, r)
	}
}
