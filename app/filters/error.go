package filters

import (
	"fmt"
	"net/http"

	"github.com/mholt/binding"
)

// RenderInputError ...
func RenderInputError(w http.ResponseWriter, r *http.Request, errs binding.Errors) {
	render := GetRenderer(r)
	var errors []string
	for _, e := range errs {
		fmt.Println("---------------")
		fmt.Println(e.Fields())
		fmt.Println("---------------")
		err := e.Fields()[0] + " " + e.Error()
		errors = append(errors, err)
	}
	render.JSON(w, http.StatusBadRequest, map[string]interface{}{
		"errors": errors,
	})
}

// RenderError ...
func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	render := GetRenderer(r)
	render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
		"errors": []string{err.Error()},
	})
}
