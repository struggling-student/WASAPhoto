package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photos []database.Photo
	var photoList database.Photos
	var dbuser database.User
	var token uint64
	// get the token from the header
	token = getToken(r.Header.Get("Authorization"))
	// get the username from the url
	username := ps.ByName("username")
	user.Username = username
	user.Id = token
	// check if the user is an existing one
	dbuser, err := rt.db.GetUserById(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	// get the list of photos from the database
	photos, err = rt.db.GetPhotos(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photoList.Identifier = token
	photoList.Photos = photos
	for i := 0; i < len(photoList.Photos); i++ {
		temp := photoList.Photos[i].File
		data, err := base64.StdEncoding.DecodeString(temp)
		if err != nil {
			log.Fatal("error:", err)
		}
		img, _, _ := image.Decode(bytes.NewReader(data))
		var filename string = "./service/database/images/" + strconv.FormatInt(int64(i), 10) + ".png"
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		err = png.Encode(out, img)
		err = jpeg.Encode(out, img, nil)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)
}
