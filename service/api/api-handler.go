package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Method for logging in the user
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User - Functions for user actions
	// Method for setting a new username
	rt.router.PUT("/user/:username/setusername", rt.wrap(rt.setMyUserName))
	// Method for getting a stream of photos for the user
	rt.router.GET("/user/:username/stream", rt.wrap(rt.getMyStream))
	// Method for getting a user profile
	rt.router.GET("/users/:username/profile", rt.wrap(rt.getUserProfile))

	// Photo - Functions for photo actions
	// Method for uploading photos
	rt.router.PUT("/users/:username/photo/:photoid", rt.wrap(rt.uploadPhoto))
	// Method for deleting photos
	rt.router.DELETE("/users/:username/photo/:photoid", rt.wrap(rt.deletePhoto))
	// Method for getting photos
	rt.router.GET("/users/:username/photo", rt.wrap(rt.getUserPhotos))

	// Ban - Functions for ban actions
	// Method for banning a user
	rt.router.PUT("/users/:username/ban/:banid", rt.wrap(rt.banUser))
	// Method for deleting bans
	rt.router.DELETE("/users/:username/ban/:banid", rt.wrap(rt.unbanUser))
	// Method for getting bans
	rt.router.GET("/users/:username/ban", rt.wrap(rt.getBans))

	// Follow - Functions for follow actions
	// Method for following a user
	rt.router.PUT("/users/:username/follow/:followid", rt.wrap(rt.followUser))
	// Method for unfollowing a user
	rt.router.DELETE("/users/:username/follow/:followid", rt.wrap(rt.unfollowUser))
	// Method for getting the followers
	rt.router.GET("/users/:username/follow", rt.wrap(rt.getFollowers))

	// Like - Functions for like actions
	// Method for liking a photo
	rt.router.PUT("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.likePhoto))
	// Method for deleting a like
	rt.router.DELETE("/users/:username/photo/:photoid/like/:likeid", rt.wrap(rt.unlikePhoto))
	// Method for getting the likes
	rt.router.GET("/users/:username/photo/:photoid/like", rt.wrap(rt.getLikes))

	// Comment - Functions for comment actions
	// Method for adding a comment to a photo
	rt.router.PUT("/users/:username/photo/:photoid/comment/:commentid", rt.wrap(rt.commentPhoto))
	// Method for deleting a comment from a photo
	rt.router.DELETE("/users/:username/photo/:photoid/comment/:commentid", rt.wrap(rt.uncommentPhoto))
	// Method for getting comments from a photo
	rt.router.GET("/users/:username/photo/:photoid/comment", rt.wrap(rt.getComments))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
