package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var ban Ban
	var follow Follow
	var user User
	var dbuser database.User
	var dbban database.Ban
	var dbfollow database.Follow
	var token uint64

	token = getToken(r.Header.Get("Authorization"))
	id, err := strconv.ParseUint(ps.ByName("banid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := ps.ByName("username")
	user.Username = username

	// check if the user is an existing one
	dbuser, err = rt.db.GetUserById(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	ban.BanId = id
	ban.BannedId = user.Id
	ban.UserId = token

	// add the ban to the database
	dbban, err = rt.db.CreateBan(ban.BanToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// remove the user from the followers database
	follow.FollowedId = user.Id
	follow.UserId = token
	dbfollow, err = rt.db.GetFollowId(follow.FollowToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	follow.FollowFromDatabase(dbfollow)
	err = rt.db.RemoveFollow(follow.FollowToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// remove all comments from the banned user on the user's posts
	err = rt.db.RemoveComments(token, user.Id)
	// remove all likes from the banned user on the user's posts
	err = rt.db.RemoveLikes(token, user.Id)
	// return the user banned to the user
	ban.BanFromDatabase(dbban)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ban)
}
