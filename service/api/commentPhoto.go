package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong JSON received")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.isValid() {
		rt.baseLogger.WithError(err).Warning("wrong user format received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment.IDPhoto = photoId
	dbComment, err := rt.db.CommentPhoto(comment.ToDatabase())

	if err != nil {
		ctx.Logger.WithError(err).Error("Comment cannot be uploaded")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UpdateCommentsPhoto(comment.IDPhoto, 1)
	if err != nil && errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be added")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.FromDatabase(dbComment)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}
