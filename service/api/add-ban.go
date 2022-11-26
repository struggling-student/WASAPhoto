package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ban Ban
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	test := r.Header.Get("Authorization")
	tokenS := re.FindAllString(test, -1)
	token, err := strconv.Atoi(tokenS[0])
	username := ps.ByName("username")
	banId, err := strconv.Atoi(ps.ByName("banid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Error := rt.db.SetBan(username, token, int(banId))
	if Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ban.Username = username
	ban.Identifier = token
	ban.BanIdentifier = banId
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ban)

}
