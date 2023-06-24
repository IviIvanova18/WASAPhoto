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
	username := ps.ByName("username")
	token, _ := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)

	dbUser, err := rt.db.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user.FromDatabase(dbUser)

	err = rt.db.IsBanned(user.UserID, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user.FollowersCount = uint64(len(dbUser.Followers))
	user.FollowingsCount = uint64(len(dbUser.Followings))
	user.PhotosCount = uint64(len(dbUser.PhotosId))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
