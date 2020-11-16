package hello

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Router       *mux.Router
	Debug        *log.Logger
	Info         *log.Logger
	Warn         *log.Logger
	Error        *log.Logger
	ShutdownChan chan os.Signal
}
type logType string

var empty = log.New(ioutil.Discard, "", log.Ldate)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) debug() *log.Logger {
	return s.log("debug")
}
func (s *Server) info() *log.Logger {
	return s.log("info")
}
func (s *Server) warn() *log.Logger {
	return s.log("warn")
}
func (s *Server) error() *log.Logger {
	return s.log("error")
}

func (s *Server) log(t logType) (l *log.Logger) {
	switch t {
	case "debug":
		l = s.Debug
	case "info":
		l = s.Info
	case "warn":
		l = s.Warn
	case "error":
		l = s.Error
	}
	if l == nil {
		return empty
	}
	return l
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	err := json.NewEncoder(w).Encode(data)
	if err == nil {
		return
	}
	s.error().Printf("Write failed: %s\n", err)
}

func (s *Server) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
