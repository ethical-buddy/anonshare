package utils

import (
	"net/http"
	"encoding/json"

	
)

// utils/json_response.go
func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		// In case of an encoding error, send fallback error response
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	}
}