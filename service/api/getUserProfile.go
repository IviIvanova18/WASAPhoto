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

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	user.Username = ps.ByName("username")
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var header = strings.Split(r.Header.Get("Authorization"), " ")
	token, _ := strconv.ParseUint(header[1], 10, 64)
	if token != userID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	dbUser, err := rt.db.GetUserProfile(user.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user.FromDatabase(dbUser)
	err = rt.db.IsBanned(user.UserID, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("Can't find banned users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user.FollowersCount = uint64(len(dbUser.Followers))
	user.FollowingsCount = uint64(len(dbUser.Followings))
	user.PhotosCount = uint64(len(dbUser.PhotosId))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
