package main

import (
	"context"
	"fmt"
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
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	log.Print("Server gracefully stopped")
}

func run() error {
	port := parseServerPort(os.Getenv(serverPortEnvKey))
	errChan := make(chan os.Signal, 1)
	signal.Notify(errChan, syscall.SIGTERM, syscall.SIGKILL)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(port), 10),
		Handler: hello.NewHandler(errChan),
	}
	go startServer(server)()
	log.Printf("Server is listening on http://localhost%s\n", server.Addr)
	<-errChan
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
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
			log.Fatalf("listen: %s\n", err)
		}
	}
}
