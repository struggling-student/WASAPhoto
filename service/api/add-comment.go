package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photo Photo
	var dbuser database.User
	var dbphoto database.Photo
	var dbcomment database.Comment
	var comment Comment
	var token uint64
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	commentid, err := strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	comment.UserId = token
	comment.Id = commentid
	comment.PhotoId = photoid
	dbcomment, err = rt.db.SetComment(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentFromDatabase(dbcomment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}
