package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rafaelzig/go-rest/internal/app/hello"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const serverPortEnvKey = "SERVER_PORT"
const defaultServerPort = uint16(8080)

func main() {
	h := createHandler()
	srv := createServer(h)
	shutdownChan := make(chan struct{})
	go handleGracefulShutdown(h, srv, shutdownChan)
	h.Info.Printf("Starting HTTP server on http://localhost%s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		h.Error.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-shutdownChan
	h.Info.Print("HTTP server gracefully Shutdown")
}

func createHandler() *hello.Server {
	h := &hello.Server{
		Router: mux.NewRouter(),
		Debug:  log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		Info:   log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warn:   log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:  log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	h.Routes()
	return h
}

func createServer(h *hello.Server) *http.Server {
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(parseServerPort(os.Getenv(serverPortEnvKey))), 10),
		Handler: h,
	}
	return server
}

func handleGracefulShutdown(h *hello.Server, srv *http.Server, shutdownChan chan struct{}) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM, syscall.SIGKILL)
	<-interruptChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Info.Println("Initiating HTTP server Shutdown")
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		h.Info.Printf("HTTP server Shutdown: %v", err)
	}
	close(shutdownChan)
}

func parseServerPort(str string) uint16 {
	if len(str) == 0 {
		return defaultServerPort
	}
	serverPort, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return defaultServerPort
	}

	return uint16(serverPort)
}
