package handler

import (
	"net/http"
)

/*Some action*/
func DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The diagnostics site.", http.StatusOK)
}
