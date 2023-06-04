package api

import (
	"encoding/json"
	"net/http"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var users []database.UserLogin

	users, err = rt.db.ListUsers()
	if err != nil {
		ctx.Logger.WithError(err).Error("can't list users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var frontendUsers = make([]UserLogin, len(users))
	for idx := range users {
		frontendUsers[idx].FromDatabase(users[idx])
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendUsers)
}
