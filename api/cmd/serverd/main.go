package main

import (
	"fmt"
	"net/http"

	"github.com/tamilchiran/nebula-api/api/internal/service/comics"
)

func main() {
	// Define the handler function for your endpoint

	// Register the handler function for a specific route
	http.HandleFunc("/article", comics.PublishArticle)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
