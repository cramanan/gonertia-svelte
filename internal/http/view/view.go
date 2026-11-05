// Package view renders Inertia pages for every handler that needs one.
package view

import (
	"log/slog"
	"net/http"

	"github.com/romsar/gonertia/v3"
)

// View renders Inertia pages. run builds the one instance and hands it to the
// handlers that render.
type View struct {
	inertia *gonertia.ViteInstance
}

func New(i *gonertia.ViteInstance) *View { return &View{inertia: i} }

// Write renders component on the response.
func (v *View) Write(w http.ResponseWriter, r *http.Request, component string, props gonertia.Props) error {
	// Exposes the page to the root template so viteAssets can preload its lazy chunk.
	r = r.WithContext(gonertia.SetTemplateDatum(r.Context(), "component", component))
	return v.inertia.Render(w, r, component, props)
}

// Page renders a component that needs no props.
func (v *View) Page(component string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := v.Write(w, r, component, nil); err != nil {
			slog.ErrorContext(r.Context(), "render", "component", component, "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})
}

// Middleware is gonertia's Inertia request handling.
func (v *View) Middleware(next http.Handler) http.Handler {
	return v.inertia.Middleware(next)
}
