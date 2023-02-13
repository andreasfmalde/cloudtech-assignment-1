package handler

import "net/http"

func DefaultPathHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The default site.", http.StatusOK)
}
