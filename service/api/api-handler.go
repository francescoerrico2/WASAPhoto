package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.sessionHandler))
	rt.router.GET("/users", rt.wrap(rt.getUsersQuery))
	rt.router.PUT("/users/:id", rt.wrap(rt.putNickname))
	rt.router.GET("/users/:id", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:id/banned_users/:banned_id", rt.wrap(rt.putBan))
	rt.router.DELETE("/users/:id/banned_users/:banned_id", rt.wrap(rt.deleteBan))
	rt.router.PUT("/users/:id/followers/:follower_id", rt.wrap(rt.putFollow))
	rt.router.DELETE("/users/:id/followers/:follower_id", rt.wrap(rt.deleteFollow))
	rt.router.GET("/users/:id/home", rt.wrap(rt.getHome))
	rt.router.POST("/users/:id/photos", rt.wrap(rt.postPhoto))
	rt.router.DELETE("/users/:id/photos/:photo_id", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:id/photos/:photo_id", rt.wrap(rt.getPhoto))
	rt.router.POST("/users/:id/photos/:photo_id/comments", rt.wrap(rt.postComment))
	rt.router.DELETE("/users/:id/photos/:photo_id/comments/:comment_id", rt.wrap(rt.deleteComment))
	rt.router.PUT("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.putLike))
	rt.router.DELETE("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.deleteLike))
	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
