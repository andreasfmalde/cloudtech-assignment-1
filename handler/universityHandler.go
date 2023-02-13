package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
	"strings"
)

func UniversityHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only "+http.MethodGet+" requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	// The fourth element is the search word
	search := strings.Split(r.URL.Path, "/")[4]

	if search == "" {
		http.Error(w, "Input a name of a university", http.StatusForbidden)
		return
	}

	var universityList []global.UniversityInformationStruct

	universityList, err := RequestUniversityInformation(search)

	if err != nil {
		http.Error(w, "Could not obtain a universitylist", http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	er := encoder.Encode(universityList)

	if er != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
	}

	//http.Error(w, "The university information site. "+r.Method, http.StatusOK)
}
