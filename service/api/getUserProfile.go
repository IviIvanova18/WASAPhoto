package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)


// getMyStream retuens a JSON list of photos
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var users = []User{
		User{
			ID: 1,
			Username : "User1",
			PhotoCount: 3,
			Followers: 12,
			Followings: 23,
		},
		User{
			ID: 2,
			Username : "User2",
			PhotoCount: 0,
			Followers: 2,
			Followings: 54,
		},User{

			ID: 3,
			Username : "User3",
			PhotoCount: 43,
			Followers: 102,
			Followings: 203,
		},

	}
	w.Header().Set("content-type", "application/json")
	err:= json.NewEncoder(w).Encode(photos)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("getUserProfile returned an error on encoding")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message:"Internal server error!"})
	}
}