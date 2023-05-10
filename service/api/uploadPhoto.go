package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"fmt"
)

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	buf, err := ioutil.ReadAll(r.Body)
	authorizationHeader := r.Header.Get("Authorization")
	fmt.Println(authorizationHeader)
	id_u, _ := strconv.ParseUint(r.Header.Get("Authorization")[7:], 10, 64)

	var photo Photo
	photo.Path = buf
	photo.IDUser = id_u
	photo.Likes = 0
	photo.Comments = 0
	photo.DateTime = time.Now()

	dbPhoto, err := rt.db.UploadPhoto(photo.ToDatabase())

	if err != nil {
		ctx.Logger.WithError(err).Error("Photo cannot be uploaded")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// err = rt.db.ChangeNoPhotoesLikesUser(id_u, 1)
	// if err != nil && errors.Is(err, database.ErrUserDoesNotExist) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// } else if err != nil {
	// 	ctx.Logger.WithError(err).WithField("id", id_u).Error("User was not found")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	photo.FromDatabase(dbPhoto)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Identifier{Id: photo.IDPhoto})
}
