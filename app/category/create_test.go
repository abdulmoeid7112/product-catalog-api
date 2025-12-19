package catalog

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/mytheresa/go-hiring-challenge/repositories/mocks"
)

// TestCategoryCreateHandler tests the category creation handler
func TestCategoryCreateHandler(t *testing.T) {
	mockRepo := &mocks.MockCategoryRepository{}
	handler := NewCategoryHandler(mockRepo)

	t.Run("creates a new category successfully", func(t *testing.T) {
		payload := map[string]string{
			"code": "ACCESSORIES",
			"name": "Accessories",
		}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.HandleCreate(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Status  bool
			Payload models.Category
		}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.True(t, resp.Status)
		assert.Equal(t, "ACCESSORIES", resp.Payload.Code)
		assert.Equal(t, "Accessories", resp.Payload.Name)
		assert.NotZero(t, resp.Payload.ID)
	})

	t.Run("returns error on invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/categories", bytes.NewBuffer([]byte(`invalid`)))
		rec := httptest.NewRecorder()

		handler.HandleCreate(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp struct {
			Status bool
			Errors any
		}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.False(t, resp.Status)
	})
}
