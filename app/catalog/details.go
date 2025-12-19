package catalog

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mytheresa/go-hiring-challenge/app/api"
)

// HandleDetail Get product details by code
func (h *CatalogHandler) HandleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	if code == "" {
		api.ErrorResponse(w, http.StatusBadRequest, productCodeRequired, nil)
		return
	}

	product, err := h.repo.GetByCode(r.Context(), code)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, productDetailsFailure, err.Error())
		return
	}

	if product == nil {
		api.ErrorResponse(w, http.StatusNotFound, productNotExist, nil)
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

	api.OKResponse(w, ProductDetailResponse{
		Code:  product.Code,
		Price: product.Price.String(),
		Category: CategoryResponse{
			Code: product.Category.Code,
			Name: product.Category.Name,
		},
		Variants: variants,
	}, productDetailsSuccess)
}
