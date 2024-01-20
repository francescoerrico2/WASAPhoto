package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userToFollowId := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	if requestingUserId == userToFollowId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ps.ByName("follower_id") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banned, err := rt.db.BannedUserCheck(
		database.User{IdUser: requestingUserId},
		database.User{IdUser: userToFollowId})
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/rt.db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.FollowUser(
		User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: userToFollowId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	oldFollower := ps.ByName("follower_id")
	photoOwnerId := ps.ByName("id")

	valid := validateRequestingUser(oldFollower, requestingUserId)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	if photoOwnerId == requestingUserId {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	banned, err := rt.db.BannedUserCheck(
		database.User{IdUser: requestingUserId},
		database.User{IdUser: photoOwnerId})
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/rt.db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.UnfollowUser(
		User{IdUser: oldFollower}.ToDatabase(),
		User{IdUser: photoOwnerId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-follow: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
