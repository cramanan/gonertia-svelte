package main

import (
	_ "embed"
	"govite/vite"
	"log"
	"net/http"

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

	// Now create your HTTP server.
	// Gonertia works well with standard http server library,
	// but you are free to use some external routers like Gorilla Mux or Chi.
	if dist := vite.Dist(); dist != nil {
		http.Handle("/build/", http.StripPrefix("/build/", http.FileServerFS(dist)))
	}
	http.Handle("/", vi.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := vi.Render(w, r, "index", inertia.Props{"key": "value"}); err != nil {
			log.Println(err)
		}
	})))

	http.ListenAndServe(":8080", nil)
}
