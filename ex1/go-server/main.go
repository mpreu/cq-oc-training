package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HelloWorldResponse holds the message
type HelloWorldResponse struct {
	Message string
}

func main() {
	http.HandleFunc("/", root)
	log.Println("Starting server on port 64912")
	if err := http.ListenAndServe(":64912", nil); err != nil {
		panic(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request '/'")
	responseBody := HelloWorldResponse{
		Message: "Hello World",
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
