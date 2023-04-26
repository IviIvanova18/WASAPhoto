package api

import (
	"database/sql"
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) GetPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	idPhoto, err := strconv.ParseUint(ps.ByName("idPhoto"), 10, 64)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't parte uint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buf, err := rt.db.GetPhoto(idPhoto)
	if errors.Is(sql.ErrNoRows, err) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Photo not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_, _ = w.Write(buf)

}
