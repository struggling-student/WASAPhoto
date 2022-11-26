package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the username from the URL
	currentUser := ps.ByName("username")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	_ = r.Body.Close()
	if err != nil {
		// the body was not a parsable JSON, reject it.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var identifier int
	identifier, err = rt.db.SetUsername(currentUser, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Identifier = identifier // return username and identifier with a 201 created response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
