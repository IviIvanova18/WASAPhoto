package api

import (
	
	"database/sql"
	"encoding/json"
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	user.Username = ps.ByName("username")
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	dbUser, err := rt.db.GetUserProfile(user.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user.FromDatabase(dbUser)
	err = rt.db.IsBanned(user.IDUser, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("Can't find banned users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user.FollowersCount = len(dbUser.Followers)
	user.FollowingsCount = len(dbUser.Followings)
	user.PhotosCount = len(dbUser.PhotosId)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
