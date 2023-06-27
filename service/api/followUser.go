package api

import (
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	idUser, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusBadRequest)
		return
	}
	if token != idUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	idFollowed, err := strconv.ParseUint(ps.ByName("followedUserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.FollowUser(idUser, idFollowed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
