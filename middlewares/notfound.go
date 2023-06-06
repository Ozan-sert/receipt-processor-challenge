package middlewares

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
	
}