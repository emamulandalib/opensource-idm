package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api", Router())

	port := 8888
	addr := fmt.Sprintf(":%d", port)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		fmt.Printf("Listening on %s\n", addr)

		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Printf("closing http server...\ns")
				return
			}
			os.Exit(1)
		}
	}()

	sig := <-signalChan
	fmt.Printf("Received signal: %q, shutting down...", sig.String())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Avoid context leak

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("could not gracefully shutdown\n")
		os.Exit(1)
	}
}
