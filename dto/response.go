package dto

import "time"

// APIResponse represents a generic API response
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Timestamp  time.Time   `json:"timestamp"`
	Pagination Pagination  `json:"pagination"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success   bool      `json:"success"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Code      string    `json:"code,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// NewAPIResponse creates a new API response
func NewAPIResponse(success bool, message string, data interface{}) APIResponse {
	return APIResponse{
		Success:   success,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(error, message, code string) ErrorResponse {
	return ErrorResponse{
		Success:   false,
		Error:     error,
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
	}
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(message string, data interface{}, page, limit, total int) PaginatedResponse {
	totalPages := (total + limit - 1) / limit
	return PaginatedResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	}
}
