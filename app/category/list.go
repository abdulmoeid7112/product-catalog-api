package catalog

import (
	"net/http"

	"github.com/mytheresa/go-hiring-challenge/app/api"
)

// HandleList List all categories
func (h *CategoryHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetAll(r.Context())
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, categoryListFailure, err.Error())
		return
	}

	// Convert to DTO
	type CategoryResponse struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	resp := make([]CategoryResponse, len(categories))
	for i, c := range categories {
		resp[i] = CategoryResponse{
			Code: c.Code,
			Name: c.Name,
		}
	}

	api.OKResponse(w, resp, categoryListSuccess)
}
