package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	photoOwnerId := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	banned, err := rt.db.BannedUserCheck(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: photoOwnerId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if len(comment.Comment) > 30 {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: comment longer than 30 characters")
		return
	}

	photo_id_64, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/ParseInt: failed convert photo_id to int64")
		return
	}

	commentId, err := rt.db.CommentPhoto(PhotoId{IdPhoto: photo_id_64}.ToDatabase(), User{IdUser: requestingUserId}.ToDatabase(), comment.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/db.CommentPhoto: failed to execute query for insertion")
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(CommentId{IdComment: commentId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/Encode: failed convert photo_id to int64")
		return
	}

}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	banned, err := rt.db.BannedUserCheck(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: ps.ByName("id")}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	photo_id_64, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
		return
	}

	comment_id_64, err := strconv.ParseInt(ps.ByName("comment_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
		return
	}

	if ps.ByName("id") == requestingUserId {

		err = rt.db.UncommentPhotoAuthor(PhotoId{IdPhoto: photo_id_64}.ToDatabase(), CommentId{IdComment: comment_id_64}.ToDatabase())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("post-comment: failed to execute query for insertion")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = rt.db.UncommentPhoto(PhotoId{IdPhoto: photo_id_64}.ToDatabase(), User{IdUser: requestingUserId}.ToDatabase(), CommentId{IdComment: comment_id_64}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment: failed to execute query for insertion")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
