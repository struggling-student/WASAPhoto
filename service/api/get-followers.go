package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var follows []database.Follow
	var followList database.Followers
	var dbuser database.User
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
	// get the bans from the database
	follows, err = rt.db.GetFollowers(user.ToDatabase())
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
