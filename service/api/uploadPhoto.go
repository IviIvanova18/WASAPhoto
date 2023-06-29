package api

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username, _ := rt.db.GetUsernameByUserID(userID)
	photo.UserID = userID
	photo.Likes = 0
	photo.Comments = 0
	photo.DateTime = time.Now()
	photo.Username = username
	// fmt.Println(photo.Path)

	token, err := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusBadRequest)
		return
	}
	if token != userID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	dbPhoto, err := rt.db.UploadPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't upload image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(dbPhoto)
	// fmt.Println(photo.Path, photo.Username, photo.DateTime)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
