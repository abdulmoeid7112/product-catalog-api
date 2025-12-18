package catalog

import (
	"net/http"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

func (h *CatalogHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	offset := 0
	if v := q.Get("offset"); v != "" {
		if o, err := strconv.Atoi(v); err == nil && o >= 0 {
			offset = o
		}
	}

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
		api.ErrorResponse(w, http.StatusInternalServerError, productListFailure, err.Error())
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

	api.OKResponse(w, response, productListSuccess)
}
