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

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var header = strings.Split(r.Header.Get("Authorization"), " ")
	token, _ := strconv.ParseUint(header[1], 10, 64)
	var idUser, _ = rt.db.GetIDByPhotoID(photoId)
	err = rt.db.IsBanned(token, idUser)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	commentsDB, err := rt.db.GetCommentsOfImage(photoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var comments []Comment
	var comment Comment

	for _, dbcom := range commentsDB {
		comment.FromDatabase(dbcom)
		comments = append(comments, comment)
	}
	if len(comments) == 0 {
		comments = []Comment{}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comments)

}
