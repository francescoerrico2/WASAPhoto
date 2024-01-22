package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pathId := ps.ByName("id")
	pathBannedId := ps.ByName("banned_id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	valid := validateRequestingUser(pathId, requestingUserId)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	if requestingUserId == pathBannedId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.BanUser(User{IdUser: pathId}.ToDatabase(), User{IdUser: pathBannedId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.BanUser: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnfollowUser(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: pathBannedId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.UnfollowUser1: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnfollowUser(
		User{IdUser: pathBannedId}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.UnfollowUser2: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerToken := extractBearer(r.Header.Get("Authorization"))
	pathId := ps.ByName("id")
	userToUnban := ps.ByName("banned_id")
	valid := validateRequestingUser(pathId, bearerToken)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	if userToUnban == bearerToken {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err := rt.db.UnbanUser(
		User{IdUser: pathId}.ToDatabase(),
		User{IdUser: userToUnban}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-ban: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
