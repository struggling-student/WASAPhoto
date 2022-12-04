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
	var follow Follow
	var user User
	var dbuser database.User
	var dbfollow database.Follow
	var token uint64

	token = getToken(r.Header.Get("Authorization"))
	id, err := strconv.ParseUint(ps.ByName("followid"), 10, 64)
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
	follow.FollowId = id
	follow.FollowedId = user.Id
	follow.UserId = token
	dbfollow, err = rt.db.SetFollow(follow.FollowToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	follow.FollowFromDatabase(dbfollow)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(follow)
}
