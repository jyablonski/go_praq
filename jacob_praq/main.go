package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func handlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/uuid", uuidHandler)

	return mux
}

func startHTTPServer() {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handlers(),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	slog.Info("Starting http server", "addr", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		slog.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}

func main() {

	go startHTTPServer()

	// Block the main function to keep the server running
	select {}
}
