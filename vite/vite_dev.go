//go:build dev

package vite

import (
	"io/fs"

	"github.com/romsar/gonertia/v3"
)

var Vite = gonertia.NewVite

// Dist returns nil in dev mode: assets are served directly by the Vite
// dev server, not by this Go process.
func Dist() fs.FS {
	return nil
}
