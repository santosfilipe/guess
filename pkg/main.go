package main

import (
	"log"
	"net/http"
	"os"

	"github.com/santosfilipe/guess/pkg/handlers"
)

func main() {
	log.Println("GuessApp is running!")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/guess/", handlers.GuessHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
