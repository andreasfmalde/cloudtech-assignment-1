package handler

import "net/http"

func UniversityHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The university information site.", http.StatusOK)
}
