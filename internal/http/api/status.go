package api

import (
	"encoding/json/v2"
	"log/slog"
	"net/http"
)

func Status() http.Handler {
	type response struct {
		Status string `json:"status"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := encode(w, http.StatusOK, response{Status: "ok"}); err != nil {
			slog.ErrorContext(r.Context(), "status", "error", err)
		}
	})
}

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.MarshalWrite(w, v)
}
