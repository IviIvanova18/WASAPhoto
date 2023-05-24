package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	// USER 
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	rt.router.GET("/users/", rt.wrap(rt.listUsers))
	rt.router.PUT("/users/:userId/", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:userId/", rt.wrap(rt.getMyUsername))

	rt.router.GET("/users/:userId/stream/", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:userId/profile/:username/", rt.wrap(rt.getUserProfile))

	// Ban 
	rt.router.PUT("/users/:userId/banned/:bannedUserId/", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userId/banned/:bannedUserId/", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:userId/banned/", rt.wrap(rt.GetAllBannedUsers))

	// Follow 
	rt.router.PUT("/users/:userId/following/:followedUserId/", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userId/following/:followedUserId/", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:userId/followings/", rt.wrap(rt.getFollowings))
		
	// Photo
	rt.router.POST("/users/:userId/photos/",rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoId/",rt.wrap(rt.deletePhoto))
	rt.router.GET("/photos/:photoId/", rt.wrap(rt.GetPhoto))

	// Comment 
	rt.router.POST("/photos/:photoId/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:photoId/comments/:commentId/", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/photos/:photoId/comments/", rt.wrap(rt.getComments))

	// Likes 
	rt.router.PUT("/photos/:photoId/likes/:userId/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:photoId/likes/:userId/", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/photos/:photoId/likes/", rt.wrap(rt.GetAllLikesPhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
