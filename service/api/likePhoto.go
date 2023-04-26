package api

import (
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	sqlite3 "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

func (rt *_router) LikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := strconv.ParseUint(r.Header.Get("Authorization")[7:], 10, 64)
	if token != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.LikePhoto(photoId, userId)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code == sqlite3.ErrConstraint &&
				sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}

	if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be liked")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UpdateLikesPhoto(photoId, 1)
	if err != nil && errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be liked")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
