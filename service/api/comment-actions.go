package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var dbuser database.User
	var dbcomment database.Comment
	var comment Comment

	// take the comment from the body
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := getToken(r.Header.Get("Authorization"))
	// get the photoid
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// create the user structure
	username := ps.ByName("username")

	// check if the user is an existing one
	dbuser, err = rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)

	commentid, err := strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	// create the comment structure
	comment.Id = commentid
	comment.UserId = token
	comment.PhotoOwner = user.Id
	comment.PhotoId = photoid
	// set the comment in the database
	dbcomment, err = rt.db.SetComment(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentFromDatabase(dbcomment)
	// return the comment
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var comment Comment

	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// get the photo id from the url
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	commentid, err := strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get the username from the url
	username := ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	comment.Id = commentid
	comment.PhotoId = photoid
	comment.UserId = token
	comment.PhotoOwner = user.Id
	dbcomment, err := rt.db.GetCommentById(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentFromDatabase(dbcomment)
	// delete the comment
	err = rt.db.RemoveComment(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var dbuser database.User
	var comments []database.Comment
	var commentList database.Comments

	token := getToken(r.Header.Get("Authorization"))
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := ps.ByName("username")
	// check if the user is an existing one
	dbuser, err = rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)

	comments, err = rt.db.GetComments(photoid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	commentList.RequestIdentifier = token
	commentList.PhotoIdentifier = photoid
	commentList.PhotoOwner = user.Id
	commentList.Comments = comments

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(commentList)
}
