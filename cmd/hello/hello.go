package main

import (
	"context"
	"fmt"
	"github.com/rafaelzig/go-rest/internal/app/hello"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const serverPortEnvKey = "SERVER_PORT"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	log.Print("Server gracefully stopped")
}

func run() error {
	errChan := make(chan os.Signal, 1)
	signal.Notify(errChan, syscall.SIGTERM, syscall.SIGKILL)
	server := hello.NewServerFacade(os.Getenv(serverPortEnvKey), errChan)
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is listening on http://localhost%s\n", server.Addr)
	<-errChan // blocks until SIGTERM or SIGKILL is received
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
