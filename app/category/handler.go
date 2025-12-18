package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/models"
)

// CategoryHandler handles category-related HTTP requests
type CategoryHandler struct {
	repo models.CategoryRepository
}

// NewCategoryHandler creates a new CategoryHandler
func NewCategoryHandler(repo models.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}
