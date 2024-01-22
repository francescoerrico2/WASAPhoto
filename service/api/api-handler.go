package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:id", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:id", rt.wrap(rt.getProfile))
	rt.router.PUT("/users/:id/banned_users/:banned_id", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:id/banned_users/:banned_id", rt.wrap(rt.unbanUser))
	rt.router.PUT("/users/:id/followers/:follower_id", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:id/followers/:follower_id", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:id/home", rt.wrap(rt.getMyStream))
	rt.router.POST("/users/:id/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:id/photos/:photo_id", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:id/photos/:photo_id", rt.wrap(rt.getPhoto))
	rt.router.POST("/users/:id/photos/:photo_id/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:id/photos/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
