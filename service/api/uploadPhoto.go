package api

import (
	"encoding/json"
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

	photo.IDUser = userID
	photo.Likes = 0
	photo.Comments = 0
	photo.DateTime = time.Now()

	dbPhoto, err := rt.db.UploadPhoto(photo.ToDatabase())
	if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", userID).Error("User is not found. Can't upload photo!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(dbPhoto)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
