package api

import (
	"encoding/json"
	"errors"

	// "fmt"
	"net/http"
	"strconv"

	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	uid, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Getting the error, no Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
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

	updatedUsername.ID = uid

	err = rt.db.SetMyUserName(updatedUsername.ToDatabase())
	if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", uid).Error("can't update the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updatedUsername)

}
