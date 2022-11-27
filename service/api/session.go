package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) session(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user User
	var dbuser database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	_ = r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	var identifier int = rand.Intn(max-min+1) + min
	user.Identifier = identifier
	dbuser, _ = rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	// return username and identifier with a 201 created response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
