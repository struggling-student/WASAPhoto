package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login - Functions for logging in
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User - Functions for user actions
	rt.router.PUT("/user/:username/setusername", rt.wrap(rt.setMyUserName))
	rt.router.GET("/user/:username/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/user/:username/profile", rt.wrap(rt.getUserProfile))

	// Photo - Functions for photo actions
	rt.router.PUT("/users/:username/photo/:photoid", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:username/photo/:photoid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:username/photo", rt.wrap(rt.getUserPhotos))

	// Ban - Functions for ban actions
	rt.router.PUT("/users/:username/ban/:banid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/ban/:banid", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:username/ban", rt.wrap(rt.getBans))

	// Follow - Functions for follow actions
	rt.router.PUT("/users/:username/follow/:followid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/follow/:followid", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:username/follow", rt.wrap(rt.getFollowers))

	// Like - Functions for like actions
	rt.router.PUT("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/users/:username/photo/:photoid/like", rt.wrap(rt.getLikes))

	// Comment - Functions for comment actions
	rt.router.PUT("/users/:username/photo/:photoid/comment/:commentid", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:username/photo/:photoid/comment/:commentid", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/users/:username/photo/:photoid/comment", rt.wrap(rt.getComments))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
