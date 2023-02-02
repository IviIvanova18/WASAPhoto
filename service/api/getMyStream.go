package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)


// getMyStream retuens a JSON list of photos
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var photos = []Photo{
		Photo{
			ID:1,
			DateTime: "01.03.2023",
			Likes:23,
			Comments:[]string{"New","Another Comment"},
		},
		Photo{
			ID:2,
			DateTime: "04.03.2022",
			Likes:2543,
			Comments:[]string{"New Com"},
		},
		Photo{
			ID:3,
			DateTime: "02.04.2023",
			Likes:232,
			Comments:[]string{"New"},
		},
	}
	w.Header().Set("content-type", "application/json")
	err:= json.NewEncoder(w).Encode(photos)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("listStreamPhotos returned an error on encoding")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message:"Internal server error!"})
	}
}
