package inertia

import (
	"net/http"
	"sync"

	"github.com/romsar/gonertia/v3"
)

var (
	instance *gonertia.ViteInstance
	once     sync.Once
)

func Init(i *gonertia.ViteInstance) { once.Do(func() { instance = i }) }

func Render(w http.ResponseWriter, r *http.Request, component string, props ...gonertia.Props) error {
	return instance.Render(w, r, component)
}

func Middleware(next http.Handler) http.Handler {
	return instance.Middleware(next)
}
