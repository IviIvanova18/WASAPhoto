package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
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

	token, _ := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if token != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.UnbanUser(userId, bannedId)
	if errors.Is(err, database.ErrBanDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", userId).Error("can't unban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
