package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	idPhoto, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idComment, err := strconv.ParseUint(ps.ByName("commentId"), 10, 64)
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

	idUser, _ := rt.db.GetIDByPhotoID(idPhoto)

	isBanned, err := rt.db.IsBanned(idUser, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err == nil && isBanned {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = rt.db.UncommentPhoto(idComment, idPhoto)
	if errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", idComment).Error("Comment cannot be deleted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UpdateCommentsPhoto(idPhoto, -1)
	if err != nil && errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Comment cannot be deleted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
