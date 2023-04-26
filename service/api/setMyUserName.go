package api

import (
	"encoding/json"
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	sqlite3 "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

func (rt *_router) SetMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := strconv.ParseUint(r.Header.Get("Authorization")[7:], 10, 64)
	if token != uid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var newUserName UserLogin
	err = json.NewDecoder(r.Body).Decode(&newUserName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !newUserName.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUserName.ID = uid
	err = rt.db.SetMyUserName(newUserName.ToDatabase())
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code == sqlite3.ErrConstraint &&
				sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {

				username, err := rt.db.GetUsernameById(newUserName.ID)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				newUserName.Username = username

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(newUserName)
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("The Username cannot be updated")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newUserName)

}
