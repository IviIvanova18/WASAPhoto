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

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)

	idUser, _ := rt.db.GetIDByPhotoID(photoId)
	err = rt.db.IsBanned(idUser, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	commentsDB, err := rt.db.GetCommentsOfImage(photoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comments := make([]Comment, len(commentsDB))
	for idx := range commentsDB {
		comments[idx].FromDatabase(commentsDB[idx])
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comments)

}
