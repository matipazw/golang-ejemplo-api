package httpHelpers

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HttpNotFound(w http.ResponseWriter) {

	var v string = "Item Not Found"
	w.WriteHeader(http.StatusNotFound)
	Json(w, v)
}
