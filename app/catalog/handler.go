package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/repositories"
)

type CatalogHandler struct {
	repo repositories.ProductRepository
}

func NewCatalogHandler(r repositories.ProductRepository) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}
