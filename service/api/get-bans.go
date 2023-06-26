package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetAllBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var listBannedUsers []database.Ban

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse uint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)

	err = rt.db.IsBanned(userId, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	

	listBannedUsers, err = rt.db.GetAllBannedUsersDB(userId)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't get listBannedUsers users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var frontendBans = make([]Ban, len(listBannedUsers))
	for idx := range listBannedUsers {
		frontendBans[idx].FromDatabase(listBannedUsers[idx])
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendBans)
}
