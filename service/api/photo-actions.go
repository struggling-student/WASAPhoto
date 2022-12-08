package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// TODO DESCRIPTION
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User
	var photo Photo

	token := getToken(r.Header.Get("Authorization"))
	user.Id = token
	user.Username = ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.CheckUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)

	// If the user is confirmed then he can upload a photo
	// Get the uploaded photo
	id, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	photo.File = encoded

	currentTime := time.Now()
	photo.Date = currentTime.Format("2006-01-02 15:04:05")

	photo.UserId = user.Id
	photo.Id = id

	dbphoto, err := rt.db.SetPhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(dbphoto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}

// TODO DESCRIPTION
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// get the photo id from the url
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get the username from the url
	username := ps.ByName("username")
	user.Username = username
	user.Id = token
	// check if the user is an existing one
	dbuser, err := rt.db.CheckUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	// delete the photo from the database
	err = rt.db.RemovePhoto(photoid)
	if errors.Is(err, database.ErrPhotoDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", photoid).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// !TESTARE POSSIIBLI ERRORI
func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photoList database.Photos

	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// get the username from the url
	username := ps.ByName("username")

	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	// get the list of photos from the database
	photos, err := rt.db.GetPhotos(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photoList.Identifier = token
	photoList.Photos = photos
	for i := 0; i < len(photoList.Photos); i++ {
		temp := photoList.Photos[i].File
		data, err := base64.StdEncoding.DecodeString(temp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var filename string = "./service/database/images/" + strconv.FormatInt(int64(i), 10) + ".png"
		out, err := os.Create(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = png.Encode(out, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = jpeg.Encode(out, img, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)
}
