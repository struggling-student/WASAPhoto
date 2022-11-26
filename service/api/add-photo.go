package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the username from the URL
	currentUser := ps.ByName("username")
	// Get the photo from the URL
	id, err := strconv.ParseInt(ps.ByName("photoid"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var photo Photo
	error := json.NewDecoder(r.Body).Decode(&photo)
	_ = r.Body.Close()
	if error != nil {
		// the body was not a parsable JSON, reject it.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo.Username = currentUser
	photo.PhotoIdentifier = id
}
