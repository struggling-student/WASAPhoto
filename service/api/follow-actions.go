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

// followUser is a function that allows a user to follow another user, it takes the username from the path and the followid from the path and returns the follow body in the response.
// It also sets the ban status to 0 (not banned).
// It returns an error if the user is not found or if the followid does not exists.
// Authorizations: the user that wants to follow another user must be logged in.
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the follow
	var follow Follow
	// struct for the user
	var user User

	// get the username from the path
	username := ps.ByName("username")
	// get the id of the username from the database
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from the database
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
	// get the bearer token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the userId from the token in the follow struct (the user that follows)
	follow.UserId = token
	// set the ban status to 0 (not banned)
	follow.BanStatus = 0

	// set the follow in the database
	dbfollow, err := rt.db.SetFollow(follow.FollowToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the follow from the database
	follow.FollowFromDatabase(dbfollow)

	// set the header and return the follow body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(follow)
}

// unfollowUser is a function that allows a user to unfollow another user, it takes the username from the path and the followid from the path and returns a response if the follow is removed from the database.
// It returns an error if the user is not found or if the followid does not exists.
// Authorizations: the user that wants to remove the follow must be logged in.
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the follow
	var follow Follow
	// struct for the user
	var user User

	// get the username from the path
	username := ps.ByName("username")
	// get the id of the username from the database
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from the database
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
	// get the bearer token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the userId from the token in the follow struct (the user that follows)
	follow.UserId = token

	// remove the follow from the database
	err = rt.db.RemoveFollow(follow.FollowId, follow.UserId, follow.FollowedId)
	if errors.Is(err, database.ErrBanDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("can't delete the follow")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set the header
	w.WriteHeader(http.StatusNoContent)
}

// getFollowers is a function that allows a user to get its followers, it takes the username from the path and returns the followers in the response.
// It returns an error if the user is not found or if the username and id are not matching.
// Authorizations: the user that wants to get the follow must be logged in.
func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the user
	var user User
	// struct for the follow list
	var follow Follow

	// get the bearer token from the header
	token := getToken(r.Header.Get("Authorization"))
	// get the username from the path
	user.Username = ps.ByName("username")

	// check if the user is an existing one (a user can only get its own followers, not the one of another user)
	dbuser, err := rt.db.CheckUserByUsername(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from the database
	user.FromDatabase(dbuser)
	// get the bans from the database
	dbfollow, err := rt.db.GetFollowers(user.ToDatabase(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	follow.FollowFromDatabase(dbfollow)
	// set the header and return the followList body
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(follow)
}
