package main

import (
	"log"
	"net/http"

	"golang-htmx-example/handlers"
)


const serverPort = ":8000"

func main() {
	// register routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/books", handlers.AddBookHandler)

	// start server
	log.Println("Starting server on port", serverPort, "...")
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
