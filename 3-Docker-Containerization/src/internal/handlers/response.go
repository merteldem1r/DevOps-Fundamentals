package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func SuccessJSON(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Failed to encode JSON response", "error", err)
	}
}
