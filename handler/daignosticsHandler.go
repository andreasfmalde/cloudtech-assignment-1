package handler

import (
	"net/http"
)

func DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The diagnostics site.", http.StatusOK)
}
