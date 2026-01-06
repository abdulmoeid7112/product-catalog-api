package mocks

import (
	"context"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/mytheresa/go-hiring-challenge/repositories"
)

// MockProductRepository implements ProductRepository for testing
type MockProductRepository struct {
	Products []models.Product // List of products for List endpoint
	Product  *models.Product  // Single product for GetByCode endpoint
	Err      error            // Optional error to simulate failures
}

// List mocks fetching products with filtering and pagination
func (m *MockProductRepository) List(ctx context.Context, filter repositories.ProductFilter) ([]models.Product, int64, error) {
	if m.Err != nil {
		return nil, 0, m.Err
	}

	// Filtering
	var filtered []models.Product
	for _, p := range m.Products {
		if filter.CategoryCode != "" && p.Category.Code != filter.CategoryCode {
			continue
		}
		if filter.MaxPrice != nil && p.Price.Cmp(*filter.MaxPrice) > 0 {
			continue
		}
		filtered = append(filtered, p)
	}

	// Pagination
	start := filter.Offset
	if start > len(filtered) {
		start = len(filtered)
	}
	end := start + filter.Limit
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], int64(len(filtered)), nil
}

// GetByCode mocks fetching a single product by code
func (m *MockProductRepository) GetByCode(ctx context.Context, code string) (*models.Product, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if m.Product == nil || m.Product.Code != code {
		return nil, nil
	}
	return m.Product, nil
}

// GetAll mocks fetching all products (used for legacy handlers if needed)
func (m *MockProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Products, nil
}
