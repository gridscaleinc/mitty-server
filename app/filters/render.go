package filters

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
)

const (
	renderContextKey = "filters/render_setup"
)

type renderSetupHandler struct {
	envName string
	next    http.Handler
}

func (p *renderSetupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	renderEngine := render.New(render.Options{
		IndentJSON: (r.Form.Get("pretty") == "true"),
	})

	r = r.WithContext(context.WithValue(r.Context(), renderContextKey, renderEngine))
	p.next.ServeHTTP(w, r)
}

// RenderSetupHandler prepares render.Render for each request.
func RenderSetupHandler(envName string, next http.Handler) http.Handler {
	return &renderSetupHandler{envName, next}
}

// GetRenderer returns a render.Render instance for each request.
func GetRenderer(r *http.Request) *render.Render {
	return r.Context().Value(renderContextKey).(*render.Render)
}
