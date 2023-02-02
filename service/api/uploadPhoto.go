package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)
var photos []Photo

// getMyStream retuens a JSON list of photos
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong JSON received")
		w.WriteHeader(http.StatusBadRequest)
		return 
	}else if photo.IsValid(){
		rt.baseLogger.WithError(err).Warning("JSON is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	photos = append(photos, photo)
	w.WriteHeader(http.StatusCreated)


	
}


