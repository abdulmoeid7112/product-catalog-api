package catalog

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

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

type CatalogHandler struct {
	repo models.ProductRepository
}

func NewCatalogHandler(r models.ProductRepository) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	// Offset
	offset := 0
	if v := q.Get("offset"); v != "" {
		if o, err := strconv.Atoi(v); err == nil && o >= 0 {
			offset = o
		}
	}

	// Limit
	limit := 10
	if v := q.Get("limit"); v != "" {
		if l, err := strconv.Atoi(v); err == nil {
			if l < 1 {
				limit = 1
			} else if l > 100 {
				limit = 100
			} else {
				limit = l
			}
		}
	}

	// Filters
	filter := models.ProductFilter{
		CategoryCode: q.Get("category"),
		Offset:       offset,
		Limit:        limit,
	}

	if v := q.Get("price_lt"); v != "" {
		if d, err := decimal.NewFromString(v); err == nil {
			filter.MaxPrice = &d
		}
	}

	products, total, err := h.repo.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respProducts := make([]ProductResponse, len(products))
	for i, p := range products {
		respProducts[i] = ProductResponse{
			Code:  p.Code,
			Price: p.Price.String(),
			Category: CategoryResponse{
				Code: p.Category.Code,
				Name: p.Category.Name,
			},
		}
	}

	response := Response{
		Total:    total,
		Products: respProducts,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
