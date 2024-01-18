package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUsersQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")

	identifier := extractBearer(r.Header.Get("Authorization"))

	if identifier == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	identificator := r.URL.Query().Get("id")

	res, err := rt.db.SearchUser(User{IdUser: identifier}.ToDatabase(), User{IdUser: identificator}.ToDatabase())
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
