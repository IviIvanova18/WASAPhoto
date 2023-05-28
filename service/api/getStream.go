package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user UserLogin
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var header = strings.Split(r.Header.Get("Authorization"), " ")
	token, _ := strconv.ParseUint(header[1], 10, 64)
	err = rt.db.IsBanned(token, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user.ID = id
	photosDB, err := rt.db.GetStreamFollowing(user.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var photos []Photo
	var photo Photo

	for _, img := range photosDB {
		photo.FromDatabase(img)
		photos = append(photos, photo)
	}

	if len(photos) == 0 {
		photos = []Photo{}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photos)

}
