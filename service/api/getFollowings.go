package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse uint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	listFollowedUsers, err := rt.db.GetFollowingsDB(userId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get listFollowedUsers users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type FollwedUsers struct {
		Followed []string `json:"followedusers"`
	}

	if len(listFollowedUsers) == 0 {
		listFollowedUsers = []string{}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(FollwedUsers{
		Followed: listFollowedUsers,
	})

}
