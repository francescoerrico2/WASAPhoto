package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoAuthor := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	pathLikeId := ps.ByName("like_id")

	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if photoAuthor == requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banned, err := rt.db.BannedUserCheck(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: photoAuthor}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if pathLikeId != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo_id_64, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error converting path param photo_id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.LikePhoto(PhotoId{IdPhoto: photo_id_64}.ToDatabase(), User{IdUser: pathLikeId}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoAuthor := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if photoAuthor == requestingUserId {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	banned, err := rt.db.BannedUserCheck(
		User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: photoAuthor}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	photoIdInt, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like/ParseInt: error converting photo_id to int64")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.UnlikePhoto(
		PhotoId{IdPhoto: photoIdInt}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like/db.UnlikePhoto: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
