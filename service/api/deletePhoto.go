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
	// The Photo ID in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.DeletePhoto(id)
	if errors.Is(err, database.ErrPhotoDoesNotExist) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		
		ctx.Logger.WithError(err).WithField("id", id).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}