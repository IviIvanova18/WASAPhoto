package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// rt.router.GET("/user/:id",rt.getUserProfile)
	// USER tested
	rt.router.POST("/session/", rt.wrap(rt.CreateUser))
	rt.router.GET("/users/", rt.wrap(rt.listUsers))
	rt.router.PUT("/users/:userId/", rt.wrap(rt.SetMyUserName))
	rt.router.GET("/users/:userId/stream/", rt.wrap(rt.GetStream))

	rt.router.GET("/users/:userId/Profile/:username/", rt.wrap(rt.SearchUser))


	//Ban tested
	rt.router.PUT("/users/:userId/banned/:bannedUserId/", rt.wrap(rt.BanUser))
	rt.router.DELETE("/users/:userId/banned/:bannedUserId/", rt.wrap(rt.UnbanUser))
	rt.router.GET("/users/:userId/banned/", rt.wrap(rt.GetAllBannedUsers))

	//Follow tested
	rt.router.PUT("/users/:userId/following/:followedUserId/", rt.wrap(rt.FollowUser))
	rt.router.DELETE("/users/:userId/following/:followedUserId/", rt.wrap(rt.UnfollowUser))
	rt.router.GET("/users/:userId/following/", rt.wrap(rt.GetFollowings))
	
	
	//Photo
	rt.router.POST("/users/:userId/photos/",rt.wrap(rt.UploadPhoto))
	rt.router.DELETE("/photos/:photoId/",rt.wrap(rt.DeletePhoto))
	
	rt.router.GET("/photos/:photoId/", rt.wrap(rt.GetPhoto))

	rt.router.POST("/photos/:photoId/comments/", rt.wrap(rt.CommentPhoto))
	rt.router.DELETE("/photos/:photoId/comments/:commentId/", rt.wrap(rt.UncommentPhoto))
	rt.router.GET("/photos/:photoId/comments/", rt.wrap(rt.GetAllCommentsPhoto))

	rt.router.PUT("/photos/:photoId/likes/:userId/", rt.wrap(rt.LikePhoto))
	rt.router.DELETE("/photos/:photoId/likes/:userId/", rt.wrap(rt.UnlikePhoto))



	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
