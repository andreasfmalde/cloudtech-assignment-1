package handler

import "net/http"

func NeighbouringUniHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "The neighbour site.", http.StatusOK)
}
