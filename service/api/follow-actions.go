package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// TODO DESCRIPTION
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// all variables
	var follow Follow
	var user User

	// get the bearer token from the header
	token := getToken(r.Header.Get("Authorization"))
	// get the username from the path
	username := ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	// get the followid from the path
	id, err := strconv.ParseUint(ps.ByName("followid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the followid from the path in the follow struct
	follow.FollowId = id
	// set the followedId from the user struct in the follow struct (the user that is followed)
	follow.FollowedId = user.Id
	// set the userId from the token in the follow struct (the user that follows)
	follow.UserId = token
	// Create the follow in the database
	follow.BanStatus = "False"
	dbfollow, err := rt.db.SetFollow(follow.FollowToDatabase())
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

// TODO DESCRIPTION
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	var user User

	token := getToken(r.Header.Get("Authorization"))
	id, err := strconv.ParseUint(ps.ByName("followid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)

	follow.FollowId = id
	follow.FollowedId = user.Id
	follow.UserId = token
	err = rt.db.RemoveFollow(follow.FollowId, follow.UserId, follow.FollowedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(follow)
}

// TODO DESCRIPTION
func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var followList database.Followers

	// create user structure for the user that wants to get the bans
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
	// get the bans from the database
	follows, err := rt.db.GetFollowers(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the bans to the user
	followList.Id = token
	followList.Followers = follows

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(followList)
}
