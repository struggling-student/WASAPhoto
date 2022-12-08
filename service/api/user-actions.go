package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// TODO DESCRIPTION
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// create user structure
	var user User
	// decode the json body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// create the user in the database
	dbuser, err := rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from database
	user.FromDatabase(dbuser)
	// return username and identifier
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// TODO DESCRIPTION
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// create user structure
	var user User
	// get the username from the url
	username := ps.ByName("username")
	// decode the json body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get the token from the authorization header
	token := getToken(r.Header.Get("Authorization"))
	// set the userid to the token
	user.Id = token
	// set the new username in the database
	dbuser, err := rt.db.SetUsername(user.ToDatabase(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from database
	user.FromDatabase(dbuser)
	// return new username and identifier
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// TODO DESCRIPTION
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// create user struct
	var user User
	// create profile struct
	var profile Profile

	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the user id to the token
	username := ps.ByName("username")
	// get the id of the user making the profile request
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from database
	user.FromDatabase(dbuser)
	// set the profile request id to the user id that requested the profile
	profile.RequestId = token
	// set the profile id to the user id that is the profile
	profile.Id = user.Id
	// set the profile username to the user username that is the profile
	profile.Username = user.Username
	// get the followers count
	followersCount, err := rt.db.GetFollowersCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the followers count of the profile
	profile.FollowersCount = followersCount
	// get the following count
	followingCount, err := rt.db.GetFollowingsCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the following count of the profile
	profile.FollowingCount = followingCount
	// get the photos count
	photoCount, err := rt.db.GetPhotosCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the photos count of the profile
	profile.PhotoCount = photoCount
	// return the profile
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// create user struct
	var user User
	//	create database photoList struct
	var photoList database.Steam

	// get the username from the url
	username := ps.ByName("username")
	// get the id of the user that wants the stream
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from database
	user.FromDatabase(dbuser)
	// get the stream of the user
	photos, err := rt.db.GetMyStream(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the id of the user that wants the stream
	photoList.Identifier = token
	// set the photos to the stream
	photoList.Photos = photos
	// create the directory for the stream
	var mainDir string = "./service/database/stream/" + strconv.FormatUint(token, 10)
	// create the directory
	err = os.Mkdir(mainDir, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// for loop to get all the photos
	for i := 0; i < len(photoList.Photos); i++ {
		// get the likes count
		likesCount, err := rt.db.GetLikesCount(photoList.Photos[i].Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// set the likes count for the photo
		photoList.Photos[i].LikeCount = likesCount
		// get the comments count
		commentsCount, err := rt.db.GetCommentsCount(photoList.Photos[i].Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// set the comments count for the photo
		photoList.Photos[i].CommentCount = commentsCount
		// get the base65 string of the photo
		file := photoList.Photos[i].File
		// decode the base64 string
		data, err := base64.StdEncoding.DecodeString(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// decode the image
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// create a directory for the user
		filename := mainDir + strconv.FormatInt(int64(i), 10) + ".png"
		photoList.Photos[i].File = filename
		out, err := os.Create(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = png.Encode(out, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = jpeg.Encode(out, img, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// return the stream
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)
}
