package api

import (
	"encoding/json"
	"strings"

	"net/http"
	"strconv"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	var updatedUsername UserLogin
	err = json.NewDecoder(r.Body).Decode(&updatedUsername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !updatedUsername.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUsername.ID = userID

	err = rt.db.SetMyUserName(updatedUsername.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updatedUsername)
}
