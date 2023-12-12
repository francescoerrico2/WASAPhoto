package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var User user
	var RequestUser user
	var Profile profile

	token := getToken(r.Header.Get("Authorization"))
	RequestUser.Id = token

	dbrequestuser, err := rt.db.CheckUserById(RequestUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	RequestUser.FromDatabase(dbrequestuser)

	username := ps.ByName("username")

	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	User.FromDatabase(dbuser)

	Profile.RequestId = token
	Profile.Id = User.Id
	Profile.Username = User.Username

	followersCount, err := rt.db.GetFollowersCount(User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Profile.FollowersCount = followersCount
	followingCount, err := rt.db.GetFollowingsCount(User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Profile.FollowingCount = followingCount
	photoCount, err := rt.db.GetPhotosCount(User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Profile.PhotoCount = photoCount

	Profile.BanStatus, err = rt.db.GetBanStatus(RequestUser.Id, User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Profile.FollowStatus, err = rt.db.GetFollowStatus(RequestUser.Id, User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Profile.CheckIfBanned, err = rt.db.CheckIfBanned(RequestUser.Id, User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Profile)
}
