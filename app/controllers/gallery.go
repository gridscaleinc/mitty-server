package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

// Gallery ...
type Gallery struct {
	ID        int    `json:"id"`
	Seq       int    `json:"seq"`
	Caption   string `json:"caption"`
	BriefInfo string `json:"briefInfo"`
	FreeText  string `json:"freeText"`
	EventID   int    `json:"eventId"`
	IslandID  int    `json:"islandId"`
}

// Content ...
type Content struct {
	Mime    string `json:"mime"`
	Name    string `json:"name"`
	LinkURL string `json:"link_url"`
	Data    []byte `json:"data"`
}

// GalleryContentParams ...
type GalleryContentParams struct {
	Gallery struct {
		ID        int    `json:"id"`
		Seq       int    `json:"seq"`
		Caption   string `json:"caption"`
		BriefInfo string `json:"briefInfo"`
		FreeText  string `json:"freeText"`
		EventID   int    `json:"eventId"`
		IslandID  int    `json:"islandId"`
	} `json:"gallery"`
	Content struct {
		Mime    string `json:"mime"`
		Name    string `json:"name"`
		LinkURL string `json:"link_url"`
		Data    []byte `json:"data"`
	} `json:"content"`
}

// // FieldMap defines parameter requirements
// func (p *GalleryContentParams) FieldMap(r *http.Request) binding.FieldMap {
// 	return binding.FieldMap{
// 		&p.Gallery: "gallery",
// 		&p.Content: "content",
// 	}
// }
//
// // FieldMap ...
// func (p *Gallery) FieldMap(r *http.Request) binding.FieldMap {
// 	return binding.FieldMap{
// 		&p.ID:        "id",
// 		&p.Seq:       "seq",
// 		&p.Caption:   "caption",
// 		&p.BriefInfo: "briefInfo",
// 		&p.FreeText:  "freeText",
// 		&p.EventID:   "eventId",
// 		&p.IslandID:  "islandId",
// 	}
// }
//
// // FieldMap ...
// func (p *Content) FieldMap(r *http.Request) binding.FieldMap {
// 	return binding.FieldMap{
// 		&p.Mime:     "mime",
// 		&p.Name:     "name",
// 		&p.LinkeURL: "link_url",
// 		&p.Data:     "data",
// 	}
// }

// PostGalleryContentHandler ...
func PostGalleryContentHandler(w http.ResponseWriter, r *http.Request) {
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
	// p := new(GalleryContentParams)
	// if errors := binding.Bind(r, p); errors != nil {
	// 	render.JSON(w, http.StatusBadRequest, map[string]interface{}{
	// 		"errors": errors,
	// 	})
	// 	return
	// }

	// decoder := json.NewDecoder(r.Body)
	// var p GalleryContentParams
	// if err := decoder.Decode(&p); err != nil {
	// 	fmt.Println(err)
	// 	render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
	// 		"errors": err,
	// 	})
	// 	return
	// }
	// defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}
	var p GalleryContentParams
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	putToS3(p.Content.Data, "hoge.png")

	gallery := new(models.Gallery)
	gallery.Seq = p.Gallery.Seq
	gallery.Caption = p.Gallery.Caption
	gallery.BriefInfo = p.Gallery.BriefInfo
	gallery.FreeText = p.Gallery.FreeText
	if err := gallery.Insert(*tx); err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	contents := new(models.Contents)
	contents.Mime = p.Content.Mime
	contents.Name = p.Content.Name
	contents.LinkURL = p.Content.LinkURL
	if err := contents.Insert(*tx); err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	output := map[string]interface{}{
		"result": "success",
	}

	// TODO
	fmt.Println(p.Gallery.EventID)
	fmt.Println(p.Gallery.IslandID)

	render.JSON(w, http.StatusCreated, output)
}

func putToS3(buf []byte, fileName string) {
	//s3Config := s.Config.S3Config
	cre := credentials.NewStaticCredentials(
		"AKIAI6WJQ2KNFSEAB4OQ",
		"61XGKqSGs6VcEDvOONaqp6zWbaINH1GEbTDw4fXI",
		"")

	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String("ap-northeast-1"),
	})

	reader := bytes.NewReader(buf)

	_, err := cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("mitty-image"),
		Key:    aws.String("/content/" + fileName),
		Body:   reader,
	})
	fmt.Println("===S3===")
	fmt.Println(err)
	fmt.Println("=======")
	if err != nil {
		log.Fatal(err)
	}
}
