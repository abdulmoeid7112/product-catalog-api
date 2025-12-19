package catalog

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/mytheresa/go-hiring-challenge/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCategoryListHandler(t *testing.T) {
	mockRepo := &mocks.MockCategoryRepository{
		Categories: []models.Category{
			{ID: 1, Code: "CLOTHING", Name: "Clothing"},
			{ID: 2, Code: "SHOES", Name: "Shoes"},
		},
	}
	handler := NewCategoryHandler(mockRepo)

	t.Run("returns all categories successfully", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/categories", nil)
		rec := httptest.NewRecorder()

		handler.HandleList(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Status  bool
			Payload []models.Category
		}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.True(t, resp.Status)
		assert.Len(t, resp.Payload, 2)
	})
}
