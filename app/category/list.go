package catalog

import (
	"encoding/json"
	"net/http"
)

// HandleList List all categories
func (h *CategoryHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
