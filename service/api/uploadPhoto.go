package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
	
)

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

	
	photo.Path = photo.Path
	photo.IDUser = userID
	photo.Likes = 0
	photo.Comments = 0
	photo.DateTime = time.Now()

	dbPhoto, err := rt.db.UploadPhoto(photo.ToDatabase())

	if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be uploaded")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// err = rt.db.ChangeNoPhotoesLikesUser(userID, 1)
	// if err != nil && errors.Is(err, database.ErrUserDoesNotExist) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// } else if err != nil {
	// 	ctx.Logger.WithError(err).WithField("id", userID).Error("User was not found")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	photo.FromDatabase(dbPhoto)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
