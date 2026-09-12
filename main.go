package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/romsar/gonertia/v3"
)

//go:embed resources/views/index.html
var rootTemplateBs []byte

//go:embed public
var publicFS embed.FS

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func main() {
	inertia := Must(gonertia.NewViteFromFS(
		Must(gonertia.NewFromBytes(rootTemplateBs)),
		publicFS,
	))

	mux := http.NewServeMux()
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServerFS(Must(fs.Sub(publicFS, "public/build")))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := inertia.Render(w, r, "Welcome"); err != nil {
			log.Println(err)
		}
	})

	http.ListenAndServe(":8000", inertia.Middleware(mux))
}
