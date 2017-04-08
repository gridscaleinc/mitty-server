package filters

import (
	"context"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/unrolled/render"
)

type contextKey string

const (
	renderContextKey contextKey = "filters/render_setup"
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

// RenderHTML ...
func RenderHTML(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	context := pongo2.Context{}
	tpl, err := pongo2.DefaultSet.FromFile(file)
	if err != nil {
		panic(err)
	}
	if data != nil {
		for k, v := range data.(map[string]interface{}) {
			context[k] = v
		}
	}
	tpl.ExecuteWriter(context, w)
}
