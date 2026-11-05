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
				_ = v.Write(&statusWriter{ResponseWriter: w, code: rec.code}, r, "Error", gonertia.Props{"status": rec.code})
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

// statusWriter sends code whatever status the render asks for: gonertia
// answers 200 (explicitly for Inertia JSON, implicitly for HTML), which turned
// every error page into a success.
type statusWriter struct {
	http.ResponseWriter
	code  int
	wrote bool
}

func (s *statusWriter) WriteHeader(int) {
	if !s.wrote {
		s.wrote = true
		s.ResponseWriter.WriteHeader(s.code)
	}
}

func (s *statusWriter) Write(b []byte) (int, error) {
	s.WriteHeader(s.code)
	return s.ResponseWriter.Write(b)
}
