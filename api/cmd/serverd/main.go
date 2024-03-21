package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tamilchiran/nebula-api/api/internal/service/comics"
)

func main() {
	// Define the handler function for your endpoint

	// Register the handler function for a specific route
	http.HandleFunc("/article", comics.PublishArticle)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	http.ListenAndServe(":"+port, nil)
}
