package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	//
	//
	//
	//
	fmt.Fprint(w, "Hello")
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Received request from %s %s\n", r.Method, r.URL.Path)
		next(w, r)

		duration := time.Since(start)
		log.Printf("Request from %s %s took %v\n", r.Method, r.URL.Path, duration)
	}
}

func main() {
	// Define the handler function for the /ping endpoint
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s %s\n", r.Method, r.URL.Path)
		fmt.Fprint(w, "pong")
	})

	http.HandleFunc("/hi", loggerMiddleware(hiHandler))
	http.HandleFunc("/hi/mid", loggerMiddleware(hiHandler))

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
