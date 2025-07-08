package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"meilisearch/handler"

	"github.com/meilisearch/meilisearch-go"
)

// RunAPIServer starts the API server
func RunAPIServer() {
	// Initialize Meilisearch client
	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey(os.Getenv("MASTER_KEY")))

	// Test connection to Meilisearch
	_, err := client.Health()
	if err != nil {
		log.Fatalf("Failed to connect to Meilisearch: %v", err)
	}
	fmt.Println("✅ Connected to Meilisearch successfully")

	// Setup routes
	mux := handler.SetupRoutes(client)

	// Apply middleware
	handler := handler.LoggingMiddleware(handler.CORSMiddleware(mux))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Starting API server on port %s\n", port)
	fmt.Println("📖 Available endpoints:")
	fmt.Println("   GET  /                    - API information")
	fmt.Println("   GET  /health              - Health check")
	fmt.Println("   GET  /api/products/search - Search products")
	fmt.Println("   GET  /api/products/       - Get product by ID")
	fmt.Println("   GET  /api/products/stats  - Get index statistics")

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
