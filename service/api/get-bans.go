package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetAllBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse uint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var header = strings.Split(r.Header.Get("Authorization"), " ")
	token, _ := strconv.ParseUint(header[1], 10, 64)
	err = rt.db.IsBanned(token, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusUnauthorized)
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
