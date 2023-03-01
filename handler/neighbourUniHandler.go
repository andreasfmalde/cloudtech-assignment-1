package handler

import (
	"net/http"
	"strconv"
	"strings"
)

func NeighbouringUniHandler(w http.ResponseWriter, r *http.Request) {

	// Make sure the method used is GET
	if r.Method != http.MethodGet {
		http.Error(w, "This site only allows "+http.MethodGet+" methods.", http.StatusMethodNotAllowed)
	}

	directories := strings.Split(r.URL.Path, "/")
	countryName := directories[4]
	universityname := directories[5]
	parameterMap := r.URL.Query()
	limit := 0
	val, ok := parameterMap["limit"]
	if ok {
		number, err := strconv.Atoi(val[0])
		if err != nil {
			http.Error(w, "Provide limit as a positive integer", http.StatusForbidden)
			return
		}
		limit = number
	}

	country, err := requestCountryInfoByName(countryName)
	country.Borders

}
