package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//? Login - Functions for logging in
	rt.router.POST("/session", rt.session)
	//TODO User - Functions for user actions
	rt.router.PUT("/user/:username/setusername", rt.setMyUserName)
	rt.router.GET("/user/:username/stream", rt.getMyStream)
	rt.router.GET("/user/:username/profile", rt.getUserProfile)
	//? Photo - Functions for photo actions
	rt.router.PUT("/users/:username/photo/:photoid", rt.banUser)
	rt.router.DELETE("/users/:username/photo/:photoid", rt.unbanUser)
	rt.router.GET("/users/:username/photo", rt.getBans)
	//? Ban - Functions for ban actions
	rt.router.PUT("/users/:username/ban/:banid", rt.followUser)
	rt.router.DELETE("/users/:username/ban/:banid", rt.unfollowUser)
	rt.router.GET("/users/:username/ban", rt.getFollowers)
	//? Follow - Functions for follow actions
	rt.router.PUT("/users/:username/follow/:followid", rt.uploadPhoto)
	rt.router.DELETE("/users/:username/follow/:followid", rt.deletePhoto)
	rt.router.GET("/users/:username/follow", rt.getUserPhotos)
	//? Like - Functions for like actions
	rt.router.PUT("/users/:username/photo/:photoid/like/:likeid", rt.likePhoto)
	rt.router.DELETE("/users/:username/photo/:photoid/like/:likeid", rt.unlikePhoto)
	rt.router.GET("/users/:username/photo/:photoid/like", rt.getLikes)
	//? Comment - Functions for comment actions
	rt.router.PUT("/users/:username/photo/:photoid/comment/:commentid", rt.commentPhoto)
	rt.router.DELETE("/users/:username/photo/:photoid/comment/:commentid", rt.uncommentPhoto)
	rt.router.GET("/users/:username/photo/:photoid/comment", rt.getComments)
	return rt.router
}
