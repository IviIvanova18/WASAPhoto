package api

import (
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bannedId, err := strconv.ParseUint(ps.ByName("bannedUserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var header = strings.Split(r.Header.Get("Authorization"), " ")
	token, _ := strconv.ParseUint(header[1], 10, 64)
	if token != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.UnbanUser(userId, bannedId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
