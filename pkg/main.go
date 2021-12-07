package main

import (
	"log"
	"net/http"
	"os"

	"github.com/santosfilipe/guess/pkg/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/guess/", handlers.GuessHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
