package handler

import "net/http"

/*
	Default path handler will just display an error message for the
	user, telling the user that the site does not exist
*/
func DefaultPathHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The site you are trying to reach does not exist.", http.StatusNotFound)
}
