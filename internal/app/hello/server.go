package hello

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type LoggerFunc func(v interface{})

type Server struct {
	Router       *mux.Router
	Debug        LoggerFunc
	Info         LoggerFunc
	Error        LoggerFunc
	ShutdownChan chan os.Signal
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("Write failed: %s\n", err)
		}
	}
}

func (s *Server) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
