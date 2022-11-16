package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Login action
	rt.router.POST("/", rt.session)
	// Session User actions
	rt.router.PUT("/setUsername")
	rt.router.GET("/steam")
	// User actions
	rt.router.GET("/profile")
	// User Ban actions
	rt.router.POST("/ban")
	rt.router.GET("/ban")
	rt.router.DELETE("/ban")
	// User Follow actions
	rt.router.POST("/follow")
	rt.router.GET("/follow")
	rt.router.DELETE("/follow")
	// Photo actions
	rt.router.POST("/photo")
	rt.router.GET("/photo")
	rt.router.DELETE("/photo")
	// Like actions
	rt.router.POST("/like")
	rt.router.GET("/like")
	rt.router.DELETE("/like")
	// Comment actions
	rt.router.POST("/comment")
	rt.router.GET("/comment")
	rt.router.PUT("/comment")
	rt.router.DELETE("/comment")

	return rt.router
}
