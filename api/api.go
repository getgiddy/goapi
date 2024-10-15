package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success code, usually 200
	Code int

	// Account balance
	Balance int64
}

// Error Response
type ErrorResponse struct {
	// Error code, usually 400
	Code int

	// Error message
	Message string
}

// Write error response
func writeError(w http.ResponseWriter, message string, code int) {
	response := ErrorResponse{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
	}
)
