package catalog

import (
	"encoding/json"
	"net/http"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
)

// HandleCreate Create a new category
func (h *CategoryHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, categoryInvalidJSON, err.Error())
		return
	}

	cat := &models.Category{
		Code: input.Code,
		Name: input.Name,
	}

	if err := h.repo.Create(r.Context(), cat); err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, categoryCreateFailure, err.Error())
		return
	}

	api.OKResponse(w, cat, categoryCreateSuccess)
}
