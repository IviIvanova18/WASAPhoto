package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/photos/", rt.getMyStream)
	rt.router.POST("/photos/",rt.uploadPhoto)
	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
