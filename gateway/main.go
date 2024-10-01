package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":8080" 
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("Failed to star http server %s", err.Error())
	}
}