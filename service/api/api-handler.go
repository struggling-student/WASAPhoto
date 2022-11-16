package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Login action
	rt.router.POST("/", rt.session)
	// Session User actions

	// User actions

	// User Ban actions

	// User Follow actions

	// Photo actions

	// Like actions

	// Comment actions

	return rt.router
}
