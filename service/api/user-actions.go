package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// doLogin is a function that allows a user to login, it takes the username from the body and it returns a user body with the identifier in the response.
// If the username is not in the database it creates a new user with a new identifier and returns it.
// Authorizations: none.
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

// setMyUserName is a function that allows a user to set his username, it takes the username from the body and it returns the user with the new username in the response body.
// It returns an error if the username is already taken.
// Authorizations: the user that wants to set a new username must be logged in.
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

// getUserProfile a function that allows a user to get a user profile, it takes the username from the url and it returns the profile in the response body.
// It returns an error if the user that requested the profile dpes not exist or the username in the path does not exist.
// Authorizations: the user that wants to get the profile of another user must be logged in.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// create user struct
	var user User
	// create request user struct
	var requestUser User
	// create profile struct
	var profile Profile

	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// set the token to the request user
	requestUser.Id = token
	// check if the request user does exist
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the user from the database
	requestUser.FromDatabase(dbrequestuser)

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

	profile.BanStatus, err = rt.db.GetBanStatus(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.FollowStatus, err = rt.db.GetFollowStatus(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the profile
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(profile)
}

// getMyStream a function that allows a user to get its stream, it takes the username from the url and it returns the stream in the response body.
// It returns an error if the user that requested the stream does not have a match between the username in the url and the bearer token.
// Authorizations: the user that wants to get the profile of another user must be logged in.
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// create user struct
	var user User
	//	create database photoList struct
	var photoList database.Steam

	// get the token from the header
	token := getToken(r.Header.Get("Authorization"))
	// get the username from the url
	username := ps.ByName("username")
	user.Id = token
	user.Username = username
	// get the id of the user that wants the stream
	dbuser, err := rt.db.CheckUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)

	// get the stream of the user
	photos, err := rt.db.GetMyStream(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the id of the user that wants the stream
	photoList.Identifier = token
	// set the photos to the stream
	photoList.Photos = photos

	// set the header and return the stream
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photoList)
}
