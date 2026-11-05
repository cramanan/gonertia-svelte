package middlewares

import (
	"io"
	"net/http"
	"slices"

	"gravel/internal/http/view"

	"github.com/romsar/gonertia/v3"
)

var errorCodes = []int{
	http.StatusInternalServerError,
	http.StatusServiceUnavailable,
	http.StatusNotFound,
	http.StatusUnauthorized,
}

// Error swallows error responses from next and renders the Inertia "Error"
// page instead.
func Error(v *view.View) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rec := &errorRecorder{ResponseWriter: w, code: http.StatusOK, writer: w}
			next.ServeHTTP(rec, r)
			if slices.Contains(errorCodes, rec.code) {
				_ = v.Write(w, r, "Error", gonertia.Props{"status": rec.code})
			}
		})
	}
}

type errorRecorder struct {
	http.ResponseWriter
	code   int
	writer io.Writer
}

func (rec *errorRecorder) WriteHeader(code int) {
	rec.code = code
	if slices.Contains(errorCodes, code) {
		rec.writer = io.Discard
		return
	}
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *errorRecorder) Write(b []byte) (int, error) {
	return rec.writer.Write(b)
}
