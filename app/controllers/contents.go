package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mholt/binding"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

// GetMyContentsHandler ...
func GetMyContentsHandler(w http.ResponseWriter, r *http.Request) {
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

	userID := filters.GetCurrentUserID(r)
	contents, err := models.GetContentsByUserID(tx, userID)
	if err != nil {
		filters.RenderError(w, r, err)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"contents": contents,
	})
}

// UploadContentsParams ...
type UploadContentsParams struct {
	Mime      string `json:"mime"`
	Name      string `json:"name"`
	Data      []byte `json:"data"`
	Thumbnail []byte `json:"thumbnail"`
}

// FieldMap defines parameter requirements
func (p *UploadContentsParams) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Mime: binding.Field{
			Form:     "mime",
			Required: false,
		},
		&p.Name: binding.Field{
			Form:     "name",
			Required: false,
		},
		&p.Data: binding.Field{
			Form:     "data",
			Required: true,
		},
		&p.Thumbnail: binding.Field{
			Form:     "thumbnail",
			Required: false,
		},
	}
}

// UploadContentsHandler ...
func UploadContentsHandler(w http.ResponseWriter, r *http.Request) {
	render := filters.GetRenderer(r)
	dbmap := helpers.GetPostgres()
	currentUserID := filters.GetCurrentUserID(r)
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}
	var p UploadContentsParams
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	fileName := TempFileName("image-", ".png")
	filePath := "https://s3-ap-northeast-1.amazonaws.com/mitty-image/content/" + fileName
	err = putToS3(p.Data, fileName)
	fmt.Println("----")
	fmt.Println(err)

	fileNameThumbnail := TempFileName("image-", "-thumbnail.png")
	filePathThumbnail := "https://s3-ap-northeast-1.amazonaws.com/mitty-image/content/" + fileNameThumbnail
	err = putToS3(p.Thumbnail, fileNameThumbnail)
	fmt.Println("----")
	fmt.Println(err)

	contents := new(models.Contents)
	contents.Mime = p.Mime
	contents.Name = p.Name
	contents.LinkURL = filePath
	contents.ThumbnailURL = filePathThumbnail
	contents.OwnerID = currentUserID
	if err = contents.Insert(*tx); err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	output := map[string]interface{}{
		"contentId": contents.ID,
	}

	render.JSON(w, http.StatusCreated, output)
}
