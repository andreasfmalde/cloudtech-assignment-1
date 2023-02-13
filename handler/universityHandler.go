package handler

import "net/http"

func UniversityHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only "+http.MethodGet+" requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "The university information site. "+r.Method, http.StatusOK)
}
