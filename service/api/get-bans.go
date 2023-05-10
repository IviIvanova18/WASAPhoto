package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) GetAllBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse uint - imageid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	bannedId, err := rt.db.GetAllBannedUsersDB(userId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get bannedId users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type BannedUsers struct {
		Banned []string `json:"bannedusers"`
	}

	if len(bannedId) == 0 {
		bannedId = []string{}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(BannedUsers{
		Banned: bannedId,
	})

}
