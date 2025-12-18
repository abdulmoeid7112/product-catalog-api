package catalog

type Response struct {
	Total    int64             `json:"total"`
	Products []ProductResponse `json:"products"`
}

type ProductResponse struct {
	Code     string           `json:"code"`
	Price    string           `json:"price"`
	Category CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProductDetailResponse struct {
	Code     string            `json:"code"`
	Price    string            `json:"price"`
	Category CategoryResponse  `json:"category"`
	Variants []VariantResponse `json:"variants"`
}

type VariantResponse struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price string `json:"price"`
}
