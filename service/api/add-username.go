package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User
	var dbuser database.User

	// get the bearer token from the header
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	Authorization := r.Header.Get("Authorization")
	stringToken := re.FindAllString(Authorization, -1)
	// token
	token, _ := strconv.Atoi(stringToken[0])

	err := json.NewDecoder(r.Body).Decode(&user)
	user.Id = uint64(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dbuser, err = rt.db.SetUsername(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
