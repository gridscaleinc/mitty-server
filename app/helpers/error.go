package helpers

import (
	"net/http"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
)

// RenderInputError ...
func RenderInputError(w http.ResponseWriter, r *http.Request, errs binding.Errors) {
	render := filters.GetRenderer(r)
	var errors []string
	for _, e := range errs {
		err := e.Fields()[0] + " " + e.Error()
		errors = append(errors, err)
	}
	render.JSON(w, http.StatusBadRequest, map[string]interface{}{
		"errors": errors,
	})
}

// RenderDBError ...
func RenderDBError(w http.ResponseWriter, r *http.Request, err error) {
	render := filters.GetRenderer(r)
	render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
		"errors": []string{err.Error()},
	})
}
