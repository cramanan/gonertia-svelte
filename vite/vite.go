//go:build !dev

package vite

import (
	"embed"
	"io/fs"

	"github.com/romsar/gonertia/v3"
)

//go:embed dist
var dist embed.FS

func Vite(i *gonertia.Inertia, opts ...gonertia.ViteOption) (*gonertia.ViteInstance, error) {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		return nil, err
	}
	opts = append([]gonertia.ViteOption{gonertia.WithBuildManifest("manifest.json")}, opts...)
	return gonertia.NewViteFromFS(i, sub, opts...)
}

// Dist returns the embedded build output filesystem for serving static
// assets (e.g. under /build/). Returns nil in dev mode, where assets are
// served directly by the Vite dev server instead.
func Dist() fs.FS {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		return nil
	}
	return sub
}
