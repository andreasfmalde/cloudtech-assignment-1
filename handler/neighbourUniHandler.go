package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func NeighbouringUniHandler(w http.ResponseWriter, r *http.Request) {

	// Make sure the method used is GET
	if r.Method != http.MethodGet {
		http.Error(w, "This site only allows "+http.MethodGet+" methods.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("content-type", "application/json")

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
	if err != nil {
		http.Error(w, "The country provided can no be found, make sure the country name are written in english", http.StatusNotFound)
		return
	}

	var neighbourCountries []string
	for _, cca3 := range country.Borders {
		c, e := requestCountryInfoByAlpha(cca3, "cca3")
		if e != nil {
			http.Error(w, "Something went wrong retrieving neighbouring countries", http.StatusInternalServerError)
			return
		}
		neighbourCountries = append(neighbourCountries, c.Name["common"].(string))
	}

	var universityList []global.UniversityInformationStruct

	for _, neighbourCountry := range neighbourCountries {
		url := global.UNIVERSITY_API_URL + "search?name=" + universityname + "&country=" + neighbourCountry
		countryUniversityList, err1 := RequestUniversityInformation(url)
		if err1 != nil {
			http.Error(w, "Could not obtain a universitylist", http.StatusInternalServerError)
			return
		}
		universityList = append(universityList, countryUniversityList...)
	}

	if limit > 0 && limit < len(universityList) {
		universityList = universityList[:limit]
	}

	displayUniversitySlice(w, universityList)
}

func displayUniversitySlice(w http.ResponseWriter, uniList []global.UniversityInformationStruct) {
	encoder := json.NewEncoder(w)

	encoderror := encoder.Encode(uniList)

	if encoderror != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
	}
}
