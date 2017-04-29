package controllers

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

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

	fileName := TempFileName("image-", ".png")
	filePath := "https://s3-ap-northeast-1.amazonaws.com/mitty-image/content/" + fileName
	err = putToS3(p.Content.Data, fileName)
	fmt.Println("----")
	fmt.Println(err)

	contents := new(models.Contents)
	contents.Mime = p.Content.Mime
	contents.Name = p.Content.Name
	contents.LinkURL = filePath //p.Content.LinkURL
	if err = contents.Insert(*tx); err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	gallery := new(models.Gallery)
	gallery.Seq = p.Gallery.Seq
	gallery.Caption = p.Gallery.Caption
	gallery.BriefInfo = p.Gallery.BriefInfo
	gallery.FreeText = p.Gallery.FreeText
	gallery.ContentID = contents.ID
	if err = gallery.Insert(*tx); err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	event, err := models.GetEventByID(tx, p.Gallery.EventID)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}
	event.GalleryID = gallery.ID
	if err = event.Update(*tx); err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	island, err := models.GetIslandByID(tx, p.Gallery.IslandID)
	if err != nil {
		fmt.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}
	island.GalleryID = gallery.ID
	if err := island.Update(*tx); err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"errors": err,
		})
		return
	}

	output := map[string]interface{}{
		"result": "success",
	}

	render.JSON(w, http.StatusCreated, output)
}

func putToS3(buf []byte, fileName string) error {
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
		ACL:    aws.String("public-read"),
		Body:   reader,
	})
	if err != nil {
		return err
	}
	return nil
}

// TempFileName generates a temporary filename for use in testing or whatever
func TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 8)
	rand.Read(randBytes)
	return prefix + hex.EncodeToString(randBytes) + suffix
}
