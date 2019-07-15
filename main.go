package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	w.WriteHeader(200)
	w.Write([]byte("Long HTTP operation done"))
}

func gracefulStop(ctx context.Context, ch chan<- struct{}, srv *http.Server) {
	gracefulStop := make(chan os.Signal, 2)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-gracefulStop

	log.Printf("Caught sig: %+v", sig)

	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown Error: %v", err)
	}
	log.Printf("HTTP server gracefully stopped")

	close(ch)
}

func cleanup() {
	log.Printf("Cleaning up...")
	// your cleanup procedures here
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var router = mux.NewRouter()
	router.HandleFunc("/", GetHandler).Methods("GET")

	var srv = &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	idleConnsClosed := make(chan struct{})
	go gracefulStop(ctx, idleConnsClosed, srv)

	defer cancel()
	defer cleanup()
	log.Printf("HTTP server started")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
