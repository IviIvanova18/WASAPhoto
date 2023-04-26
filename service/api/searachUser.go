package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	user.Username = ps.ByName("username")
	userID, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := strconv.ParseUint(r.Header.Get("Authorization")[7:], 10, 64)
	if token != userID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	dbUser, err := rt.db.SearchUser(user.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user.FromDatabase(dbUser)

	// err = rt.db.IsBanned(user.IDUser, userID)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	ctx.Logger.WithError(err).Error("can't see bans")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// } else if err == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
