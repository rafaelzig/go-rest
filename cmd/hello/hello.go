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

var debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
var infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
var errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	if err := run(); err != nil {
		errorLogger.Fatalf("%s\n", err)
	}
	infoLogger.Print("Server gracefully stopped")
}

func run() error {
	h := createHandler()
	defer close(h.ShutdownChan)
	signal.Notify(h.ShutdownChan, syscall.SIGTERM, syscall.SIGKILL)
	srv := createServer(h)
	go startServer(srv)()
	infoLogger.Printf("Server listening on http://localhost%s\n", srv.Addr)
	<-h.ShutdownChan
	infoLogger.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

func createHandler() *hello.Server {
	h := &hello.Server{
		Router:       mux.NewRouter(),
		ShutdownChan: make(chan os.Signal, 1),
		Debug:        debugLogger,
		Info:         infoLogger,
		Warn:         warnLogger,
		Error:        errorLogger,
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

func startServer(server *http.Server) func() {
	return func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			errorLogger.Fatalf("listen: %s\n", err)
		}
	}
}
