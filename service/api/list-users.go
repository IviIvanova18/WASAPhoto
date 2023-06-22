package api

import (
	"encoding/json"
	"net/http"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	users, err := rt.db.ListUsers()
	if err != nil {
		ctx.Logger.WithError(err).Error("can't list users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	frontendUsers := make([]UserLogin, len(users))
	for idx := range users {
		frontendUsers[idx].FromDatabase(users[idx])
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendUsers)
}
