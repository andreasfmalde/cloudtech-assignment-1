package handler

import (
	"assignment-1/global"
	"assignment-1/handler/request"
	"encoding/json"
	"net/http"
	"strings"
)

func UniversityHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only "+http.MethodGet+" requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("content-type", "application/json")

	// The fourth element is the search word
	search := strings.Split(r.URL.Path, "/")[4]

	if search == "" {
		http.Error(w, "Input a name of a university", http.StatusForbidden)
		return
	}

	var universityList []global.UniversityInformationStruct

	search = strings.Replace(search, " ", "%20", -1)

	url := global.UNIVERSITY_API_URL + "search?name=" + search

	universityList, err := request.RequestUniversityInformation(url)

	if err != nil {
		http.Error(w, "Could not obtain a universitylist", http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	encoderror := encoder.Encode(universityList)

	if encoderror != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
	}

	//http.Error(w, "The university information site. "+r.Method, http.StatusOK)
}
