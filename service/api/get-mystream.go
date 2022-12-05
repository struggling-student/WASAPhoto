package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photos []database.PhotoStream
	var photoList database.Steam
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
	photos, err = rt.db.GetMyStream(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photoList.Identifier = token
	photoList.Photos = photos
	var mainDir string = "./service/database/stream/" + strconv.FormatUint(token, 10)
	err = os.Mkdir(mainDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(photoList.Photos); i++ {
		var commentsCount int
		var likesCount int
		// count all the likes
		likesCount, err = rt.db.GetLikesCount(photoList.Photos[i].Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		photoList.Photos[i].LikeCount = likesCount
		// count all the comments
		commentsCount, err = rt.db.GetCommentsCount(photoList.Photos[i].Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		photoList.Photos[i].CommentCount = commentsCount
		temp := photoList.Photos[i].File
		data, err := base64.StdEncoding.DecodeString(temp)
		if err != nil {
			log.Fatal("error:", err)
		}
		img, _, _ := image.Decode(bytes.NewReader(data))
		// create a directory for the user
		var filename string = mainDir + strconv.FormatInt(int64(i), 10) + ".png"
		photoList.Photos[i].File = filename
		out, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = png.Encode(out, img)
		err = jpeg.Encode(out, img, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)

}
