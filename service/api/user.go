package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization")) //prendiamo id utente
	valid := validateRequestingUser(ps.ByName("id"), identifier)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	followers, err := rt.db.GetFollowing(User{IdUser: identifier}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var photos []database.Photo
	for _, follower := range followers {
		followerPhoto, err := rt.db.GetPhotosList(
			User{IdUserd: identifier}.ToDatabase(),
			User{IdUser: follower.IdUser}.ToDatabase())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for i, photo := range followerPhoto {
			if i >= database.PhotosPerUserHome {
				break
			}
			photos = append(photos, photo)
		}
	}
	w.WriteHeader(http.StatusOk)

	_ = json.NewEncoder(w).Encode(photos)

}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	requestesUser := ps.ByName("id")

	var followers []database.User
	var following []database.User
	var photos []database.Photo

	userBanned, err := rt.db.BannedUserCheck(User{UserId: requestingUserId}.ToDatabase(), User{UserId: requestedUser}.ToDatabase)
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BannedUserCheck/userBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	requestedProfileBanned, err := rt.db.BannedUserCheck(User{IdUser: requestedUser}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BannedUserCheck/requestedProfileBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestedProfileBanned {
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	userExists, err := rt.db.CheckUser(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.CheckUser: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !userExists {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	followers, err = rt.db.GetFollowers(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetFollowers: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err = rt.db.GetFollowing(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetFollowing: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photos, err = rt.db.GetPhotosList(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetPhotosList: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	nickname, err := rt.db.GetUsername(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetNickname: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(CompleteProfile{
		Name:      requestedUser,
		Nickname:  nickname,
		Followers: followers,
		Following: following,
		Posts:     photos,
	})

}

func (rt *_router) getUsersQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization"))
	if identifier == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	identificator := r.Url.Query().Get("id")
	res, err := rt.db.SearchUser(User{UserId: identifier}.ToDatabase(), User{UserId: identificator}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Database has encountered an error")
		_ = json.NewEncoder(w).Encode([]User{})
		return
	}
	w.WriteHeader(http.StatusOK)
	if len(res) == 0 {
		_ = json.NewEncoder(w).Encode([]User{})
		return
	}
	_ = json.NewEncoder(w).Encode(res)

}

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	pathId := ps.ByName("id")
	valid := validateRequestingUser(pathId, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	var Username newUsername
	err := json.NewDecoder(r.Body).Decode(&Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-nickname: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.ModifyUsername(
		User{IdUser: pathId}.ToDatabase(),
		nick.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("update-nickname: error executing update query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
