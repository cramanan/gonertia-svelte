package main

import (
	"embed"
	"encoding/json/v2"
	"gravel/inertia"
	"io"
	"io/fs"
	"log"
	"net/http"
	"slices"

	"github.com/romsar/gonertia/v3"
)

//go:embed resources/views/index.html
var rootTemplateBs []byte

//go:embed public
var publicFS embed.FS

func main() {
	// This sets up inertia as a package
	inertia.Init(Must(gonertia.NewViteFromFS(
		Must(gonertia.NewFromBytes(rootTemplateBs)),
		publicFS,
	)))

	mux := http.NewServeMux()
	mux.HandleFunc("/demos", func(w http.ResponseWriter, r *http.Request) {
		inertia.Render(w, r, "Demos")
	})
	mux.HandleFunc("/demos/api-fetching", func(w http.ResponseWriter, r *http.Request) {
		inertia.Render(w, r, "demos/ApiFetching")
	})
	mux.HandleFunc("/api/status", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.MarshalWrite(w, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		inertia.Render(w, r, "Welcome")
	})
	mux.Handle("/", http.FileServerFS(Must(fs.Sub(publicFS, "public"))))

	port := ":8000"
	log.Printf("HTTP server listening on port %s\n", port)
	http.ListenAndServe(port, inertia.Middleware(ErrorMiddleware(mux)))
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

type ErrorRecorder struct {
	http.ResponseWriter
	Code   int
	Writer io.Writer
}

func NewErrorRecorder(w http.ResponseWriter) *ErrorRecorder {
	return &ErrorRecorder{ResponseWriter: w, Code: http.StatusOK, Writer: w}
}

func (recorder *ErrorRecorder) WriteHeader(code int) {
	recorder.Code = code
	if slices.Contains(ErrorCodes, code) {
		recorder.Writer = io.Discard
		return
	}
	recorder.ResponseWriter.WriteHeader(code)
}

func (recorder *ErrorRecorder) Write(b []byte) (int, error) {
	return recorder.Writer.Write(b)
}

var ErrorCodes = []int{
	http.StatusInternalServerError,
	http.StatusServiceUnavailable,
	http.StatusNotFound,
	http.StatusUnauthorized,
}

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := NewErrorRecorder(w)
		next.ServeHTTP(recorder, r)
		if slices.Contains(ErrorCodes, recorder.Code) {
			inertia.Render(w, r, "Error", gonertia.Props{"status": recorder.Code})
		}
	})
}
