package api

import (
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	

	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.DeletePhoto(id)
	if errors.Is(err, database.ErrPhotoDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Photo cannot be deleted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.DeleteLikes(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Likes of the photo cannot be deleted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.DeleteComments(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Comments of the photo cannot be deleted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var id_user, _ = rt.db.GetIDByPhotoID(id)
	err = rt.db.UpdatePhotoCountUser(id_user, -1)
	if err != nil && errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("User not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}