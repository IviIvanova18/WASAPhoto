package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// rt.router.GET("/user/:id",rt.getUserProfile)
	rt.router.POST("/session", rt.wrap(rt.CreateUser))
	rt.router.PUT("/User/:userId/username", rt.wrap(rt.SetMyUserName))
	rt.router.GET("/Users/:userId/Profile/:username", rt.wrap(rt.SearchUser))

	//Photo
	rt.router.POST("/users/:userId/photos",rt.wrap(rt.UploadPhoto))
	rt.router.DELETE("/photos/:photoId",rt.wrap(rt.DeletePhoto))
	rt.router.GET("/photos/:photoId", rt.wrap(rt.GetPhoto))

	rt.router.POST("/photos/:photoId/comments", rt.wrap(rt.CommentPhoto))
	rt.router.DELETE("/photos/:photoId/comments/:commentId", rt.wrap(rt.UncommentPhoto))
	rt.router.PUT("/photos/:photoId/likes/:userId", rt.wrap(rt.LikePhoto))
	rt.router.DELETE("/photos/:photoId/likes/:userId", rt.wrap(rt.UnlikePhoto))

	

	// rt.router.GET("/users/:id/photos", rt.getMyStream)
	
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
