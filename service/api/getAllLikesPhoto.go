package api

import (
	"encoding/json"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) GetAllLikesPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	likesDB, err := rt.db.GetAllLikesOfPhoto(photoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var likes []Like
	var like Like
	

	for _, dbcom := range likesDB {

		like.FromDatabase(dbcom)
		likes = append(likes, like)
	}
	if len(likes) == 0 {
		likes = []Like{}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(likes)

}






