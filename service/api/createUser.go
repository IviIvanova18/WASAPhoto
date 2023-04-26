package api

import (
	"encoding/json"
	"errors"
	"git.wasaphoto.ivi/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	sqlite3 "github.com/mattn/go-sqlite3"
	"net/http"
)

func (rt *_router) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong JSON received")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.isValid() {
		rt.baseLogger.WithError(err).Warning("wrong user format received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dbuser, err := rt.db.CreateUser(user.Username)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code == sqlite3.ErrConstraint &&
				sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {

				id, err := rt.db.GetIDByUsername(user.Username)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				user.ID = id

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(user)
				return
			}
		}

		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", user.Username).Error("can't retrieve user")
		return
	}
	user.FromDatabase(dbuser)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
