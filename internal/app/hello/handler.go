package hello

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"syscall"
)

type handler struct {
	router  *http.ServeMux
	errChan chan<- os.Signal
}

func NewHandler(errChan chan<- os.Signal) *handler {
	s := &handler{
		router:  http.NewServeMux(),
		errChan: errChan,
	}
	s.initRoutes()
	return s
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *handler) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
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

func (s *handler) handleHealth() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{
		Status: "ready",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}

func (s *handler) handleIndex() func(http.ResponseWriter, *http.Request) {
	response := struct {
		Status string `json:"message"`
	}{
		Status: "Hello World",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response, http.StatusOK)
	}
}

func (s *handler) handleShutdown() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.errChan <- syscall.SIGTERM
	}
}
