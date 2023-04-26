package api

import (
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) UnlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	idUser, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := strconv.ParseUint(r.Header.Get("Authorization")[7:], 10, 64)
	if token != idUser {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	idPhoto, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UnlikePhoto(idPhoto, idUser)

	if errors.Is(database.ErrLikeNotFound, err) {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be unliked")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UpdateLikesPhoto(idPhoto, -1)
	if err != nil && errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be unliked")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
