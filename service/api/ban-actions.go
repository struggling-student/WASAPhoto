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
	// create the ban structure
	ban.BanId = id
	ban.BannedId = user.Id
	ban.UserId = token
	// add the ban to the database
	dbban, err := rt.db.CreateBan(ban.BanToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ban.BanFromDatabase(dbban)

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

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var ban Ban
	var user User
	var dbuser database.User
	var token uint64

	token = getToken(r.Header.Get("Authorization"))
	id, err := strconv.ParseUint(ps.ByName("banid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := ps.ByName("username")
	user.Username = username
	// check if the user is an existing one
	dbuser, err = rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	ban.BanId = id
	ban.UserId = token
	ban.BannedId = user.Id
	err = rt.db.RemoveBan(ban.BanToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = rt.db.UpdateBanStatus(0, user.Id, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ban)
}

func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
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

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(banList)

}
