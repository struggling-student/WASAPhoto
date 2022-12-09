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
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the ban
	var ban Ban
	// struct for the user
	var user User

	// token of the user that wants to ban the other user
	token := getToken(r.Header.Get("Authorization"))
	// username of the user that has to be banned
	username := ps.ByName("username")
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// returns the struct of the user that has to be banned
	user.FromDatabase(dbuser)

	// ban id for the ban
	id, err := strconv.ParseUint(ps.ByName("banid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the ban id
	ban.BanId = id
	// set the banned user id
	ban.BannedId = user.Id
	// set the user id to the token
	ban.UserId = token

	// add the ban to the database
	dbban, err := rt.db.CreateBan(ban.BanToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the ban from the database
	ban.BanFromDatabase(dbban)

	// update the ban status to 1 (banned)
	err = rt.db.UpdateBanStatus(1, user.Id, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// remove all comments from the banned user on the user's posts
	err = rt.db.RemoveComments(token, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// remove all likes from the banned user on the user's posts
	err = rt.db.RemoveLikes(token, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the user banned to the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ban)
}

// unfollowUser is a function that allows a user to unfollow another user, it takes the username from the path and the followid from the path and returns a response if the follow is removed from the database.
// It returns an error if the user is not found or if the followid does not exists.
// Authorizations: the user that wants to remove the follow must be logged in.
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the ban
	var ban Ban
	// struct for the user
	var user User

	// get the token of the user that wants to unban the other user
	token := getToken(r.Header.Get("Authorization"))
	// get the ban id from the url
	id, err := strconv.ParseUint(ps.ByName("banid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get the username of the user that has to be unbanned
	username := ps.ByName("username")
	// set the username of the user that has to be unbanned
	user.Username = username
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from the database
	user.FromDatabase(dbuser)

	// set the ban id
	ban.BanId = id
	// set the user id to the token
	ban.UserId = token
	// set the banned user id
	ban.BannedId = user.Id

	// remove the ban from the database
	err = rt.db.RemoveBan(ban.BanToDatabase())
	if errors.Is(err, database.ErrBanDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// update the ban status to 0 (not banned)
	err = rt.db.UpdateBanStatus(0, user.Id, token)
	if errors.Is(err, database.ErrBanDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set the header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// getFollowers is a function that allows a user to get its followers, it takes the username from the path and returns the followers in the response.
// It returns an error if the user is not found or if the username and id are not matching.
// Authorizations: the user that wants to get the follow must be logged in.
func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// struct for the user
	var user User
	// struct for the ban list
	var banList database.Bans

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
	// decode the user from the database
	user.FromDatabase(dbuser)

	// get the bans from the database
	bans, err := rt.db.GetBans(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	// return the bans to the user
	banList.Identifier = user.Id
	banList.Username = user.Username
	banList.Bans = bans

	// set the header and return the ban list
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(banList)

}
