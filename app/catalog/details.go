package catalog

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *CatalogHandler) HandleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	if code == "" {
		http.Error(w, "product code required", http.StatusBadRequest)
		return
	}

	product, err := h.repo.GetByCode(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	var variants []VariantResponse
	for _, v := range product.Variants {
		price := v.Price
		if price.IsZero() {
			price = product.Price
		}
		variants = append(variants, VariantResponse{
			Name:  v.Name,
			SKU:   v.SKU,
			Price: price.String(),
		})
	}

	resp := ProductDetailResponse{
		Code:  product.Code,
		Price: product.Price.String(),
		Category: CategoryResponse{
			Code: product.Category.Code,
			Name: product.Category.Name,
		},
		Variants: variants,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
