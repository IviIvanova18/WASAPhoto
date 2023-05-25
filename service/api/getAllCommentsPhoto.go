package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
