package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/repositories"
)

// CategoryHandler handles category-related HTTP requests
type CategoryHandler struct {
	repo repositories.CategoryRepository
}

// NewCategoryHandler creates a new CategoryHandler
func NewCategoryHandler(repo repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}
