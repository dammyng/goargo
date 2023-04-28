package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	envValue := os.Getenv("ENVIRONMENT")
	if envValue == "" {
		fmt.Fprint(w, "ENVIRONMENT variable is not set")
	} else {
		fmt.Fprintf(w, "ENVIRONMENT is set to %s", envValue)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
