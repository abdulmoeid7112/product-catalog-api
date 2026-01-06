package api

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status      bool   `json:"status"`                // true for success, false for error
	Description string `json:"description,omitempty"` // optional human-readable message
	Payload     any    `json:"payload,omitempty"`     // main data
	Errors      any    `json:"errors,omitempty"`      // error details
}

// OKResponse returns a success response with optional payload
func OKResponse(w http.ResponseWriter, payload any, description string) {
	resp := ApiResponse{
		Status:      true,
		Description: description,
		Payload:     payload,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// ErrorResponse returns a failure response with HTTP status code and error details
func ErrorResponse(w http.ResponseWriter, status int, description string, errors any) {
	resp := ApiResponse{
		Status:      false,
		Description: description,
		Errors:      errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
