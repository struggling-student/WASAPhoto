package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photo Photo
	var likes []database.Like
	var likeList database.Likes
	var dbuser database.User
	var dbphoto database.Photo
	var token uint64

	// create user structure for the user that wants to get the bans
	token = getToken(r.Header.Get("Authorization"))
	user.Id = token
	user.Username = ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserById(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	photo.Id = photoid
	photo.UserId = user.Id
	// check if the photo is an existing one
	dbphoto, err = rt.db.GetPhotoById(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(dbphoto)
	likes, err = rt.db.GetLikes(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the likes
	likeList.PhotoIdentifier = photoid
	likeList.UserIdentifier = token
	likeList.Likes = likes

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(likeList)
}
