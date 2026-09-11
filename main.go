package main

import (
	_ "embed"
	"encoding/json/v2"
	"govite/vite"
	"log"
	"net/http"
	"strings"

	inertia "github.com/romsar/gonertia/v3"
)

//go:embed resources/views/index.html
var rootTemplateBs []byte

func main() {
	i, err := inertia.NewFromBytes(rootTemplateBs)
	if err != nil {
		log.Fatal(err)
	}
	vi, err := vite.Vite(i)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.MarshalWrite(w, map[string]string{"hello": "world"})
	})

	// Now create your HTTP server.
	// Gonertia works well with standard http server library,
	// but you are free to use some external routers like Gorilla Mux or Chi.
	if dist := vite.Dist(); dist != nil {
		fs := http.FileServerFS(dist)
		http.Handle("/build/", http.StripPrefix("/build/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "" || strings.HasSuffix(r.URL.Path, "/") {
				http.NotFound(w, r)
				return
			}
			fs.ServeHTTP(w, r)
		})))
	}

	http.Handle("/", vi.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if err := vi.Render(w, r, "index", inertia.Props{"key": "value"}); err != nil {
			log.Println(err)
		}

	})))

	http.ListenAndServe(":8080", nil)
}
