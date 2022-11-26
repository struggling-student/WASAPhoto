package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) session(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// If the user does not exist, it will be created, and an identifier
	// will be returned. If the user already exists, the user identifier is returned.

	// create a new user
	var user User
	// decode the body into a User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	_ = r.Body.Close()
	if err != nil {
		// the body was not a parsable JSON, reject it.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// generate a new idetifier
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	var identifier int = rand.Intn(max-min+1) + min
	// set a new username
	err = rt.db.SetUser(user.Username, identifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Identifier = identifier
	// return username and identifier with a 201 created response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
