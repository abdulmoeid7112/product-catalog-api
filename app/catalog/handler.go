package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/models"
)

type CatalogHandler struct {
	repo models.ProductRepository
}

func NewCatalogHandler(r models.ProductRepository) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}
