package routes

import (
	"net/http"

	"gravel/internal/http/api"
	"gravel/internal/http/view"
)

func Routes(mux *http.ServeMux, v *view.View, static http.Handler) {
	mux.Handle("GET /{$}", v.Page("Welcome"))
	mux.Handle("GET /demos", v.Page("Demos"))
	mux.Handle("GET /demos/api-fetching", v.Page("demos/ApiFetching"))
	mux.Handle("GET /api/status", api.Status())
	mux.Handle("/", static)
}
