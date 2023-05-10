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
		ctx.Logger.WithError(err).Error("can't parse uint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	listBannedUsers, err := rt.db.GetAllBannedUsersDB(userId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get listBannedUsers users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type BannedUsers struct {
		Banned []string `json:"bannedusers"`
	}

	if len(listBannedUsers) == 0 {
		listBannedUsers = []string{}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(BannedUsers{
		Banned: listBannedUsers,
	})

}
