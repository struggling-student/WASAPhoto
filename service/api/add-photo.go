package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var token uint64
	var user User
	var photo Photo
	var dbphoto database.Photo
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
	content, _ := ioutil.ReadAll(r.Body)
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	// data, err := base64.StdEncoding.DecodeString(encoded)
	// if err != nil {
	// 	log.Fatal("error:", err)
	// }
	// img, _, _ := image.Decode(bytes.NewReader(data))

	// //save the imgByte to file
	// out, err := os.Create("./service/database/QRImg.png")

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// err = png.Encode(out, img)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// log.Fatalf("Photo uploaded")
	//binary.Read(r.Body, binary.LittleEndian, &photo.File)
	photo.File = encoded
	currentTime := time.Now()
	photo.Date = currentTime.Format("01-02-2006")
	photo.UserId = user.Id
	photo.Id = id
	dbphoto, err = rt.db.SetPhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(dbphoto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}
