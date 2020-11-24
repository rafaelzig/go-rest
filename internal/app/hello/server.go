package hello

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router         *mux.Router
	LogHandlerFunc func(l Level, v ...interface{})
}

type Level string

const (
	TRACE Level = "trace"
	DEBUG Level = "debug"
	INFO  Level = "info"
	WARN  Level = "warn"
	ERROR Level = "error"
	FATAL Level = "fatal"
)
const supportedContentType = "application/json; charset=UTF-8"

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	header := w.Header()
	header.Set("Content-Type", supportedContentType)
	w.WriteHeader(status)
	if data == nil {
		return
	}
	err := s.encode(w, data)
	if err != nil {
		s.Error(fmt.Sprintf("Write failed: %s", err))
	}
}

func (s *Server) Trace(v ...interface{}) {
	s.log(TRACE, v...)
}
func (s *Server) Debug(v ...interface{}) {
	s.log(DEBUG, v...)
}
func (s *Server) Info(v ...interface{}) {
	s.log(INFO, v...)
}
func (s *Server) Warn(v ...interface{}) {
	s.log(WARN, v...)
}
func (s *Server) Error(v ...interface{}) {
	s.log(ERROR, v...)
}
func (s *Server) Fatal(v ...interface{}) {
	s.log(FATAL, v...)
}

func (s *Server) log(l Level, v ...interface{}) {
	if s.LogHandlerFunc == nil {
		return
	}
	s.LogHandlerFunc(l, v...)
}
