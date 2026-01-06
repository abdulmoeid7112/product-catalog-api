package catalog

// Response DTOs for product-related responses
type Response struct {
	Total    int64             `json:"total"`
	Products []ProductResponse `json:"products"`
}

// ProductResponse represents a product in the response
type ProductResponse struct {
	Code     string           `json:"code"`
	Price    string           `json:"price"`
	Category CategoryResponse `json:"category"`
}

// CategoryResponse represents a category in the response
type CategoryResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ProductDetailResponse represents detailed information about a product
type ProductDetailResponse struct {
	Code     string            `json:"code"`
	Price    string            `json:"price"`
	Category CategoryResponse  `json:"category"`
	Variants []VariantResponse `json:"variants"`
}

// VariantResponse represents a product variant in the response
type VariantResponse struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price string `json:"price"`
}
