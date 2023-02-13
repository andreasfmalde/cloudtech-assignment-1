package handler

import "net/http"

func DefaultPathHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The site you are trying to reach does not exist.", http.StatusNotFound)
}
