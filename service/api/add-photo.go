package api

import (
	"encoding/binary"
	"log"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var token uint64
	var user User
	var photo Photo

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
	// If the user is confirmed then he can upload a photo
	// Get the uploaded photo
	id, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	binary.Read(r.Body, binary.LittleEndian, &photo.File)
	currentTime := time.Now()
	photo.Date = currentTime.Format("01-02-2006")
	photo.UserId = user.Id
	photo.Id = id
	log.Fatal(photo)
}
