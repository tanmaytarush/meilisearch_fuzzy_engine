package handler

import (
	"net/http"
)

// SetupRoutes configures all the HTTP routes for the application
func SetupRoutes(client interface{}) *http.ServeMux {
	mux := http.NewServeMux()

	// Create handlers
	productHandler := NewProductHandler(client)

	// Product routes
	mux.HandleFunc("/api/products/search", productHandler.SearchProducts)
	mux.HandleFunc("/api/products/", productHandler.GetProductByID)
	mux.HandleFunc("/api/products/stats", productHandler.GetIndexStats)

	// Health check
	mux.HandleFunc("/health", productHandler.HealthCheck)

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"message": "Meilisearch Product Catalog API",
			"version": "1.0.0",
			"endpoints": {
				"search": "/api/products/search?q=<query>",
				"product": "/api/products?id=<id>",
				"stats": "/api/products/stats",
				"health": "/health"
			}
		}`))
	})

	return mux
}

// CORSMiddleware adds CORS headers to responses
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		println("Request:", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
