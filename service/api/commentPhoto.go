package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// There is no error since it is already checked in the Athorization header
	idUser, _ := strconv.ParseUint(r.Header.Get("Authorization")[7:], 10, 64)

	idPhoto, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// idPosted, err := rt.db.GetUserIDByPhoto(idPhoto)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", idPhoto).Error("Photo not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// err = rt.db.IsBanned(idPosted, idUser)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	var comment Comment

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment.IDUser = idUser
	comment.IDPhoto = idPhoto

	err = rt.db.CommentPhoto(comment.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Comment cannot be added")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UpdateCommentsPhoto(comment.IDPhoto, 1)
	if err != nil && errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Comment cannot be added")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
