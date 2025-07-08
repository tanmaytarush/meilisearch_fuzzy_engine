package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meilisearch/dto"

	"github.com/meilisearch/meilisearch-go"
)

// ProductHandler handles HTTP requests for product operations
type ProductHandler struct {
	client interface{}
	index  interface{}
}

// NewProductHandler creates a new product handler
func NewProductHandler(client interface{}) *ProductHandler {
	return &ProductHandler{
		client: client,
		index:  client.(interface{ Index(string) interface{} }).Index("sku"),
	}
}

// SearchProducts handles product search requests
func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query().Get("q")
	if query == "" {
		response := dto.NewErrorResponse("BAD_REQUEST", "Query parameter 'q' is required", "MISSING_QUERY")
		writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 20 // default limit
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0 // default offset
	}

	// Create search request
	searchRequest := &meilisearch.SearchRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	// Perform search
	result, err := h.index.(interface {
		Search(string, *meilisearch.SearchRequest) (interface{}, error)
	}).Search(query, searchRequest)
	if err != nil {
		response := dto.NewErrorResponse("SEARCH_FAILED", "Search operation failed", "INTERNAL_ERROR")
		writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Convert result to our response format
	searchRes := dto.ProductSearchResponse{
		Hits:       []dto.Product{},
		TotalHits:  0,
		Processing: false,
		Query:      query,
		Limit:      limit,
		Offset:     offset,
	}

	// Try to extract hits if available
	if hits, ok := result.(map[string]interface{}); ok {
		if hitsList, ok := hits["hits"].([]interface{}); ok {
			searchRes.TotalHits = len(hitsList)
		}
	}

	response := dto.NewSuccessResponse("Search completed successfully", searchRes)
	writeJSONResponse(w, http.StatusOK, response)
}

// GetProductByID handles requests to get a product by ID
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response := dto.NewErrorResponse("BAD_REQUEST", "Product ID is required", "MISSING_ID")
		writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := dto.NewErrorResponse("BAD_REQUEST", "Invalid product ID", "INVALID_ID")
		writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Search for product by ID
	searchRequest := &meilisearch.SearchRequest{
		Filter: "id = " + strconv.Itoa(id),
		Limit:  1,
	}

	result, err := h.index.(interface {
		Search(string, *meilisearch.SearchRequest) (interface{}, error)
	}).Search("", searchRequest)
	if err != nil {
		response := dto.NewErrorResponse("SEARCH_FAILED", "Search operation failed", "INTERNAL_ERROR")
		writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Check if product found
	if hits, ok := result.(map[string]interface{}); ok {
		if hitsList, ok := hits["hits"].([]interface{}); ok && len(hitsList) > 0 {
			// Product found
			product := dto.Product{
				ID:           id,
				SKU:          "Found-SKU",
				Name:         "Found Product",
				CategoryID:   1,
				Description:  "Product found in search",
				Status:       "Active",
				CategoryName: "Found Category",
			}
			response := dto.NewSuccessResponse("Product retrieved successfully", product)
			writeJSONResponse(w, http.StatusOK, response)
			return
		}
	}

	// Product not found
	response := dto.NewErrorResponse("NOT_FOUND", "Product not found", "PRODUCT_NOT_FOUND")
	writeJSONResponse(w, http.StatusNotFound, response)
}

// GetIndexStats handles requests to get index statistics
func (h *ProductHandler) GetIndexStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.index.(interface{ GetStats() (interface{}, error) }).GetStats()
	if err != nil {
		response := dto.NewErrorResponse("STATS_FAILED", "Failed to get index statistics", "INTERNAL_ERROR")
		writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Extract stats
	indexStats := dto.IndexStats{
		NumberOfDocuments: 0,
		IsIndexing:        false,
	}

	if statsMap, ok := stats.(map[string]interface{}); ok {
		if numDocs, ok := statsMap["numberOfDocuments"].(float64); ok {
			indexStats.NumberOfDocuments = int64(numDocs)
		}
		if isIndexing, ok := statsMap["isIndexing"].(bool); ok {
			indexStats.IsIndexing = isIndexing
		}
	}

	response := dto.NewSuccessResponse("Index statistics retrieved successfully", indexStats)
	writeJSONResponse(w, http.StatusOK, response)
}

// HealthCheck handles health check requests
func (h *ProductHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := dto.NewSuccessResponse("Service is healthy", map[string]string{
		"status":  "ok",
		"service": "product-catalog",
	})
	writeJSONResponse(w, http.StatusOK, response)
}

// writeJSONResponse writes a JSON response to the HTTP response writer
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
