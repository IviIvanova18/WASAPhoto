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

// GetAllBannedUsers is a handler that retrieves a list of banned users.
func (rt *_router) GetAllBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userIDStr := ps.ByName("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	token, err := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusBadRequest)
		return
	}

	isBanned, err := rt.db.IsBanned(userID, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err == nil && isBanned {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	listBannedUsers, err := rt.db.GetAllBannedUsersDB(userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	frontendBans := make([]Ban, len(listBannedUsers))
	for idx := range listBannedUsers {
		frontendBans[idx].FromDatabase(listBannedUsers[idx])
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendBans)
}
