package dto

import "time"

// Product represents a product in the catalog
type Product struct {
	ID                   int       `json:"id"`
	SKU                  string    `json:"sku"`
	Name                 string    `json:"name"`
	CategoryID           int       `json:"category_id"`
	Description          string    `json:"description"`
	ImageURLs            []string  `json:"image_urls"`
	MRP                  *string   `json:"mrp"`
	Status               string    `json:"status"`
	CreatedBy            int       `json:"created_by"`
	UpdatedBy            int       `json:"updated_by"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	PerUnitMRPPrice      *string   `json:"per_unit_mrp_price"`
	UnitType             *string   `json:"unit_type"`
	PerUnitSellingPrice  *string   `json:"per_unit_selling_price"`
	UnitValue            *string   `json:"unit_value"`
	SellingPrice         *string   `json:"selling_price"`
	CategoryBrandIndexID *string   `json:"category_brand_index_id"`
	IsActive             int       `json:"is_active"`
	Discount             *string   `json:"discount"`
	CategoryName         string    `json:"category_name"`
}

// ProductSearchRequest represents a search request for products
type ProductSearchRequest struct {
	Query     string `json:"q"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Category  string `json:"category,omitempty"`
	Status    string `json:"status,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}

// ProductSearchResponse represents a search response for products
type ProductSearchResponse struct {
	Hits       []Product `json:"hits"`
	TotalHits  int       `json:"total_hits"`
	Processing bool      `json:"processing"`
	Query      string    `json:"query"`
	Limit      int       `json:"limit"`
	Offset     int       `json:"offset"`
}

// ProductCreateRequest represents a request to create a new product
type ProductCreateRequest struct {
	SKU                  string   `json:"sku" validate:"required"`
	Name                 string   `json:"name" validate:"required"`
	CategoryID           int      `json:"category_id" validate:"required"`
	Description          string   `json:"description"`
	ImageURLs            []string `json:"image_urls"`
	MRP                  *string  `json:"mrp"`
	Status               string   `json:"status" validate:"required"`
	PerUnitMRPPrice      *string  `json:"per_unit_mrp_price"`
	UnitType             *string  `json:"unit_type"`
	PerUnitSellingPrice  *string  `json:"per_unit_selling_price"`
	UnitValue            *string  `json:"unit_value"`
	SellingPrice         *string  `json:"selling_price"`
	CategoryBrandIndexID *string  `json:"category_brand_index_id"`
	Discount             *string  `json:"discount"`
	CategoryName         string   `json:"category_name" validate:"required"`
}

// ProductUpdateRequest represents a request to update an existing product
type ProductUpdateRequest struct {
	ID                   int      `json:"id" validate:"required"`
	SKU                  string   `json:"sku"`
	Name                 string   `json:"name"`
	CategoryID           int      `json:"category_id"`
	Description          string   `json:"description"`
	ImageURLs            []string `json:"image_urls"`
	MRP                  *string  `json:"mrp"`
	Status               string   `json:"status"`
	PerUnitMRPPrice      *string  `json:"per_unit_mrp_price"`
	UnitType             *string  `json:"unit_type"`
	PerUnitSellingPrice  *string  `json:"per_unit_selling_price"`
	UnitValue            *string  `json:"unit_value"`
	SellingPrice         *string  `json:"selling_price"`
	CategoryBrandIndexID *string  `json:"category_brand_index_id"`
	Discount             *string  `json:"discount"`
	CategoryName         string   `json:"category_name"`
}

// IndexStats represents statistics about the Meilisearch index
type IndexStats struct {
	NumberOfDocuments int64 `json:"number_of_documents"`
	IsIndexing        bool  `json:"is_indexing"`
}

// TaskStatus represents the status of a Meilisearch task
type TaskStatus struct {
	TaskUID int64  `json:"task_uid"`
	Status  string `json:"status"`
	Error   *struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Type    string `json:"type"`
		Link    string `json:"link"`
	} `json:"error,omitempty"`
}
