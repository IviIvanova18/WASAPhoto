package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// rt.router.GET("/user/:id",rt.getUserProfile)
	rt.router.POST("/session", rt.wrap(rt.CreateUser))
	rt.router.PUT("/User/:uid/Username", rt.wrap(rt.SetMyUserName))
	rt.router.GET("/Users/:uid/Profile/:username", rt.wrap(rt.SearchUser))

	// rt.router.GET("/users/:id/photos", rt.getMyStream)
	// rt.router.POST("/users/:id/photos",rt.uploadPhoto)
	// rt.router.DELETE("/photos/:id",rt.deletePhoto)
	// rt.router.PUT("user/:id",rt.setMyUserName)
	// rt.router.GET("/users/:id/following",rt.getUserFollowing)
	// rt.router.DELETE("/users/:id/following/:id", rt.unfollowUser)
	// rt.router.PUT("/user/:idfollowing/:id", rt.followUser)
	// rt.router.DELETE("/users/:id/banned/:id", rt.unbanUser)
	// rt.router.PUT("/users/:id/banned/:id", rt.banUser)
    // rt.router.PUT("/photos/:id/likes/:id",rt.likePhoto)
    // rt.router.DELETE("/photos/:id/likes/:id",rt.unlikePhoto)
	// rt.router.POST("/photos/:id/comments",rt.commentPhoto)
	// rt.router.DELETE("/photos/:id/comments/:id",rt.uncommentPhoto)
	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
