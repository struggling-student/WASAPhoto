package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var token uint64
	var user User
	var dbuser database.User
	var photo Photo
	var dbphoto database.Photo
	// get the token from the header
	token = getToken(r.Header.Get("Authorization"))
	// get the photo id from the url
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get the username from the url
	username := ps.ByName("username")
	user.Username = username
	user.Id = token
	// check if the user is an existing one
	dbuser, err = rt.db.GetUserById(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	photo.Id = photoid
	photo.UserId = user.Id
	// check if the photo is an existing one
	dbphoto, err = rt.db.GetPhotoById(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(dbphoto)
	// delete the photo from the database
	err = rt.db.RemovePhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// photo deleted from the database
	photo.PhotoFromDatabase(dbphoto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}
