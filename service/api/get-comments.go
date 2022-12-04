package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photo Photo
	var dbuser database.User
	var dbphoto database.Photo
	var comments []database.Comment
	var commentList database.Comments
	var token uint64

	token = getToken(r.Header.Get("Authorization"))
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	// check if the photo is an existing one
	photo.Id = photoid
	photo.UserId = user.Id
	dbphoto, err = rt.db.GetPhotoById(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(dbphoto)
	comments, err = rt.db.GetComments(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	commentList.Id = token
	commentList.PhotoIdentifier = photoid
	commentList.Comments = comments

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(commentList)
}
