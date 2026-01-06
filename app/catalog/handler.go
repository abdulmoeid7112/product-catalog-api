package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/repositories"
)

// CatalogHandler handles catalog-related HTTP requests
type CatalogHandler struct {
	repo repositories.ProductRepository
}

// NewCatalogHandler creates a new CatalogHandler
func NewCatalogHandler(r repositories.ProductRepository) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}
