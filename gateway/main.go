package main

import (
	"log"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
	"github.com/fantom0052/market/common"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":3000")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("Failed to star http server %s", err.Error())
	}
}
