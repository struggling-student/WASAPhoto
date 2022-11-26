package api

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user User

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	test := r.Header.Get("Authorization")
	tokenS := re.FindAllString(test, -1)
	token, _ := strconv.Atoi(tokenS[0])

	username := ps.ByName("username")

	_, err := rt.db.GetUser(username)
	log.Fatalf("id: %v", user)

	if token != token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, Error := rt.db.GetBans(token)
	if Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
