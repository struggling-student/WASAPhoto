package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// commentPhoto is a function that allows a user to comment a photo, it takes the username,photo,comment from the path and returns the comment body in the response.
// It returns an error if the user is not found or if the photoid/commentid does not exists.
// Authorizations: the user that wants to follow another user must be logged in.
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the user
	var user User
	// struct for the comment
	var comment Comment

	// take the comment from the body
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the photoid
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
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
	// decode the user from database
	user.FromDatabase(dbuser)

	// get the comment id
	commentid, err := strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the comment id to the commentid from the url
	comment.Id = commentid
	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the user id to the token
	comment.UserId = token
	// set the photo id to the photoid from the url
	comment.PhotoId = photoid
	// set the photo owner to the user id
	comment.PhotoOwner = user.Id

	// set the comment in the database
	dbcomment, err := rt.db.SetComment(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the comment from database
	comment.CommentFromDatabase(dbcomment)

	// set the header and return the comment
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

// uncommentPhoto is a function that allows a user to remove a comment from a picture, it takes the username,photoid,commentid from the path and returns a response if the comment is removed from the database.
// It returns an error if the user is not found or if the  photoid/commentid does not exists.
// Authorizations: the user that wants to remove the follow must be logged in.
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the user
	var user User
	// struct for the comment
	var comment Comment

	// get the photo id from the url
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the comment id from the url
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
	// decode the user from database
	user.FromDatabase(dbuser)

	// set the comment id to the commentid from the url
	comment.Id = commentid
	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the user id to the token
	comment.UserId = token
	// set the photo id to the photoid from the url
	comment.PhotoId = photoid
	// set the photo owner to the user id
	comment.PhotoOwner = user.Id

	// remove the comment from the database
	err = rt.db.RemoveComment(comment.CommentToDatabase())
	if errors.Is(err, database.ErrCommentDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", commentid).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write the header
	w.WriteHeader(http.StatusNoContent)
}

// getComments is a function that allows a user to get all comments from a picture, it takes the username,photoid from the path and returns a commentList body
// It returns an error if the user is not found or if the  photoid does not exists.
// Authorizations: the user that wants to remove the follow must be logged in.
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the user
	var user User
	// struct for the request user
	var requestUser User
	// struct for the photo
	var photo Photo
	// struct for the commentList
	var commentList database.Comments

	// get the token from the header
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

	// get the photo id from the url
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if the photo is an existing one
	photo.Id = photoid
	dbphoto, err := rt.db.CheckPhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the photo from the database
	photo.PhotoFromDatabase(dbphoto)

	// check if the user from teh path is an existing one
	username := ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from database
	user.FromDatabase(dbuser)

	// get the comments from the database
	comments, err := rt.db.GetComments(photo.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the request user id to the commentList
	commentList.RequestIdentifier = requestUser.Id
	// set the photo id to the commentList
	commentList.PhotoIdentifier = photo.Id
	// set the photo owner to the commentList
	commentList.PhotoOwner = user.Id
	// set the comments to the commentList
	commentList.Comments = comments

	// set the header and return the commentList
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(commentList)
}
