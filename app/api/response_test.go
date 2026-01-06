package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestOKResponse tests the OKResponse function for successful JSON responses
func TestOKResponse(t *testing.T) {
	type samplePayload struct {
		Message string `json:"message"`
	}

	sample := samplePayload{Message: "Success"}

	t.Run("successful http200 json response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		OKResponse(recorder, sample, "Operation completed successfully")

		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code 200 OK")
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"), "Expected Content-Type to be application/json")

		expected := `{
			"status": true,
			"description": "Operation completed successfully",
			"payload": {"message":"Success"}
		}`

		assert.JSONEq(t, expected, recorder.Body.String(), "Response body does not match expected")
	})
}

// TestErrorResponse tests the ErrorResponse function for error JSON responses
func TestErrorResponse(t *testing.T) {
	t.Run("json response for a given http status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ErrorResponse(recorder, http.StatusInternalServerError, "Some error occurred", "error details")

		assert.Equal(t, http.StatusInternalServerError, recorder.Code, "Expected status code 500 Internal Server Error")
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"), "Expected Content-Type to be application/json")

		expected := `{
			"status": false,
			"description": "Some error occurred",
			"errors": "error details"
		}`

		assert.JSONEq(t, expected, recorder.Body.String(), "Response body does not match expected")
	})
}
