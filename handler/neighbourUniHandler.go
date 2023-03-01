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
	// Add content-type json to the response header
	w.Header().Add("content-type", "application/json")

	directories := strings.Split(r.URL.Path, "/") // Split directories URL path
	countryName := directories[4]                 // Name of base country from URL path
	universityname := directories[5]              // Name of univeristy from URL path
	parameterMap := r.URL.Query()                 // URL parameters
	limit := 0
	// Check if the limit parameter is provided
	val, ok := parameterMap["limit"]
	if ok {
		number, err := strconv.Atoi(val[0])
		if err != nil { // Make sure paramter is a valid integer
			http.Error(w, "Provide limit as a positive integer", http.StatusForbidden)
			return
		}
		limit = number // Set limit to parameter number if it's in range
	}
	// Retrieve neighbour countries names
	neighbouringCountries, ok := retriveNeighbours(w, countryName)
	if !ok {
		return // Terminate function if an error has occured
	}
	// Make a list of universities from the neighbour countries
	universityList, ok := makeUniversityList(w, neighbouringCountries, universityname)
	if !ok {
		return // Terminate function if an error has occured
	}
	// Show a reduced list if limit is specified
	if limit > 0 && limit < len(universityList) {
		universityList = universityList[:limit]
	}
	// Send the universitylist as JSON to the client
	encoder := json.NewEncoder(w)
	encoderror := encoder.Encode(universityList)
	// Display an error to the user if the encoding fails
	if encoderror != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
	}
}

func makeUniversityList(w http.ResponseWriter, nC []string, uN string) ([]global.UniversityInformationStruct, bool) {
	var universityList []global.UniversityInformationStruct
	// Collect universities from neighboring countries with the university name uN
	for _, neighbourCountry := range nC {
		url := global.UNIVERSITY_API_URL + "search?name=" + uN + "&country=" + neighbourCountry
		// Collect universities from each neighbouring country
		countryUniversityList, err1 := RequestUniversityInformation(url)
		if err1 != nil {
			http.Error(w, "Could not obtain a universitylist", http.StatusInternalServerError)
			return nil, false
		}
		// Append all universities to one list/slice
		universityList = append(universityList, countryUniversityList...)
	}
	return universityList, true // Return the universitylist
}

func retriveNeighbours(w http.ResponseWriter, countryName string) ([]string, bool) {
	// Retrieve the country information of the base country
	country, err := requestCountryInfo(countryName, global.NAME_TYPE)
	if err != nil {
		http.Error(w, "The country provided can no be found, make sure the country name are written in english", http.StatusNotFound)
		return nil, false
	}
	var neighbourCountries []string
	// Collect the common names of the base countrys neigbhouring countries
	for _, cca3 := range country.Borders {
		c, e := requestCountryInfo(cca3, global.CCA3_TYPE)
		if e != nil { // Error retrieveing a neighbour country
			http.Error(w, "Something went wrong retrieving neighbouring countries", http.StatusInternalServerError)
			return nil, false
		} // Append each neighbour country to the list
		neighbourCountries = append(neighbourCountries, c.Name["common"].(string))
	}
	return neighbourCountries, true // Return the list of neighbouring countires
}
