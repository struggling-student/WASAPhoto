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

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// create user structure
	var user User
	// create database user struct
	var dbuser database.User
	// decode the json body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// create the user in the database
	dbuser, err = rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from database
	user.FromDatabase(dbuser)
	// return username and identifier with a 201 created response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User
	var dbuser database.User
	var token uint64
	// get the token from the header
	token = getToken(r.Header.Get("Authorization"))

	err := json.NewDecoder(r.Body).Decode(&user)
	user.Id = token
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

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var dbuser database.User
	var profile Profile

	var token uint64
	var followersCount int
	var followingCount int
	var photoCount int

	// create user structure for the user that wants to get the bans
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

	profile.Id = user.Id
	profile.Username = user.Username
	// get the followers count
	followersCount, err = rt.db.GetFollowersCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowersCount = followersCount
	// get the following count
	followingCount, err = rt.db.GetFollowingsCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowingCount = followingCount
	// get the photos count
	photoCount, err = rt.db.GetPhotosCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.PhotoCount = photoCount

	// response
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}

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
