package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// uploadPhoto is a function that allows a user to upload a photo, it takes the username,photo,phtooid from the path and returns the photo body in the response.
// It returns an error if the user is not found or if the photoid alreaady exists.
// Authorizations: the user that wants to follow another user must be logged in.
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

	photo.File, err = io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

// deltePhoto is a function that allows a user to remove a  photo, it takes the username,photoid from the path and returns a response.
// It returns an error if the user is not found or if the photoid does not exists.
// Authorizations: the user that wants to follow another user must be logged in.
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

	// remove comments from the photo

	// remove the likes from the photo

}

// commentPhoto is a function that allows a user to get the photos of another user it takes the username from the path and returns the photolist body in the response.
// It returns an error if the user is not found.
// Authorizations: the user that wants to follow another user must be logged in.
func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var photoList database.Photos

	// create user structure for the user that wants to get the bans
	token := getToken(r.Header.Get("Authorization"))
	// set the token to the request user
	requestUser.Id = token
	// check if the request user does exist
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from the database
	requestUser.FromDatabase(dbrequestuser)

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
	photoList.RequestUser = requestUser.Id
	photoList.Identifier = token
	photoList.Photos = photos
	w.Header().Set("Content-Type", "image/*")
	_ = json.NewEncoder(w).Encode(photoList)
}
