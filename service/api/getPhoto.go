package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) GetPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	idPhoto, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse uint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoDB, err := rt.db.GetPhoto(idPhoto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photo.FromDatabase(photoDB)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)

}
