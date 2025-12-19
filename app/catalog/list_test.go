package catalog

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/mytheresa/go-hiring-challenge/repositories/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCatalogListHandler(t *testing.T) {
	products := []models.Product{
		{
			Code:  "PROD001",
			Price: decimal.NewFromFloat(10.99),
			Category: models.Category{
				Code: "CLOTHING", Name: "Clothing",
			},
		},
		{
			Code:  "PROD002",
			Price: decimal.NewFromFloat(20.0),
			Category: models.Category{
				Code: "SHOES", Name: "Shoes",
			},
		},
		{
			Code:  "PROD003",
			Price: decimal.NewFromFloat(5.0),
			Category: models.Category{
				Code: "ACCESSORIES", Name: "Accessories",
			},
		},
	}

	mockRepo := &mocks.MockProductRepository{Products: products}
	handler := &CatalogHandler{repo: mockRepo}

	t.Run("returns products with default pagination", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog", nil)
		rec := httptest.NewRecorder()

		handler.HandleList(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"status":true`)
		assert.Contains(t, rec.Body.String(), `"PROD001"`)
		assert.Contains(t, rec.Body.String(), `"total":3`)
	})

	t.Run("applies offset and limit", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog?offset=1&limit=1", nil)
		rec := httptest.NewRecorder()

		handler.HandleList(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"status":true`)
		assert.Contains(t, rec.Body.String(), `"PROD002"`)
		assert.NotContains(t, rec.Body.String(), `"PROD001"`)
	})

	t.Run("filters by category", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog?category=CLOTHING", nil)
		rec := httptest.NewRecorder()

		handler.HandleList(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"PROD001"`)
		assert.NotContains(t, rec.Body.String(), `"PROD002"`)
	})

	t.Run("filters by max price", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog?price_lt=15", nil)
		rec := httptest.NewRecorder()

		handler.HandleList(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"PROD001"`)
		assert.Contains(t, rec.Body.String(), `"PROD003"`)
		assert.NotContains(t, rec.Body.String(), `"PROD002"`)
	})

	t.Run("returns 500 if repository error", func(t *testing.T) {
		errorRepo := &mocks.MockProductRepository{Err: assert.AnError}
		handler := &CatalogHandler{repo: errorRepo}

		req := httptest.NewRequest("GET", "/catalog", nil)
		rec := httptest.NewRecorder()

		handler.HandleList(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), `"status":false`)
	})
}
