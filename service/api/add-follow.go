package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// all variables
	var follow Follow
	var user User
	var dbuser database.User
	var dbfollow database.Follow
	var token uint64

	// get the bearer token from the header
	token = getToken(r.Header.Get("Authorization"))
	// get the username from the path
	username := ps.ByName("username")
	// set the username in the user struct
	user.Username = username
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserById(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	// get the followid from the path
	id, err := strconv.ParseUint(ps.ByName("followid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// set the followid from the path in the follow struct
	follow.FollowId = id
	// set the followedId from the user struct in the follow struct (the user that is followed)
	follow.FollowedId = user.Id
	// set the userId from the token in the follow struct (the user that follows)
	follow.UserId = token
	// Create the follow in the database
	dbfollow, err = rt.db.SetFollow(follow.FollowToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	follow.FollowFromDatabase(dbfollow)
	// set the header and return the follow body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(follow)
}
