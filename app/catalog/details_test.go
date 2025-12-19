package catalog

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/mytheresa/go-hiring-challenge/repositories/mocks"
)

// TestCatalogDetailHandler tests the CatalogHandler's HandleDetail method
func TestCatalogDetailHandler(t *testing.T) {
	product := &models.Product{
		Code:  "PROD001",
		Price: decimal.NewFromFloat(10.99),
		Category: models.Category{
			Code: "CLOTHING",
			Name: "Clothing",
		},
		Variants: []models.Variant{
			{Name: "Variant A", SKU: "SKU001A", Price: decimal.NewFromFloat(11.99)},
			{Name: "Variant B", SKU: "SKU001B"}, // Should inherit product price
		},
	}

	mockRepo := &mocks.MockProductRepository{Product: product}
	handler := &CatalogHandler{repo: mockRepo}

	t.Run("returns product details with variants successfully", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog/PROD001", nil)
		rec := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/catalog/{code}", handler.HandleDetail)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		expectedVariantPrice := "10.99"
		assert.Contains(t, rec.Body.String(), `"sku":"SKU001B"`)
		assert.Contains(t, rec.Body.String(), `"price":"11.99"`)
		assert.Contains(t, rec.Body.String(), `"price":"`+expectedVariantPrice+`"`) // inherited
	})

	t.Run("returns 404 if product not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog/UNKNOWN", nil)
		rec := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/catalog/{code}", handler.HandleDetail)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), `"status":false`)
	})

	t.Run("returns 500 on repository error", func(t *testing.T) {
		errorRepo := &mocks.MockProductRepository{Err: assert.AnError}
		handler := &CatalogHandler{repo: errorRepo}

		req := httptest.NewRequest("GET", "/catalog/PROD001", nil)
		rec := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/catalog/{code}", handler.HandleDetail)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), `"status":false`)
	})
}
