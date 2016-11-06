package controllers

import (
	"net/http"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"

	"github.com/mholt/binding"
)

// SignUpParams ...
type SignUpParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// FieldMap defines parameter requirements
func (p *SignUpParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.UserName: binding.Field{
			Form:     "user_name",
			Required: true,
		},
		&p.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

// SignUpHandler ...
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	tx, err := dbmap.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	p := new(SignUpParams)
	if errors := binding.Bind(r, p); errors != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}
	user := new(models.User)
	user.UserName = p.UserName
	user.Password = p.Password
	err = user.Insert(*tx)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err,
		})
		return
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"user": user,
	})
}

// SignInHandler ...
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	tx, err := dbmap.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	p := new(SignUpParams)
	if errors := binding.Bind(r, p); errors != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}
	user := new(models.User)
	user.UserName = p.UserName
	user.Password = p.Password
	err = user.Insert(*tx)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"err": err,
		})
		return
	}
	render.JSON(w, http.StatusCreated, map[string]interface{}{
		"user": user,
	})
}
