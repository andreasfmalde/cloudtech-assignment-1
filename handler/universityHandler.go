package handler

import (
	"assignment-1/global"
	"assignment-1/handler/request"
	"encoding/json"
	"net/http"
	"strings"
)

/*
The university handler is responsible for all the
operations to be done when the /uniinfo path is reached.
The university handler will take a university name that
should also be provided in the URL and returns a JSON
list of all universities with parts of their name beeing
the same at the name provided in the URL.
*/
func UniversityHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure the method used is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Only "+http.MethodGet+" requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	// Add content-type json to the response header
	w.Header().Add("content-type", "application/json")

	// The fifth element is the search word
	search := strings.Split(r.URL.Path, "/")[4]
	// The search word can not be empty
	if search == "" {
		http.Error(w, "Input a name of a university", http.StatusForbidden)
		return
	}
	// Trim away the spaces of the search word
	search = strings.Replace(search, " ", "%20", -1)
	// Define the url for the university API
	url := global.UNIVERSITY_API_URL + "search?name=" + search

	// Request a list of universities with useful information
	var universityList []global.UniversityInformationStruct
	universityList, err := request.RequestUniversityInformation(url)
	// Handle any error that may occur in the request
	if err != nil {
		http.Error(w, "Could not obtain a universitylist", http.StatusInternalServerError)
	}

	// Send the universitylist as JSON to the client
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t") // Make JSON prettier to read
	encoderror := encoder.Encode(universityList)
	// Display an error to the user if the encoding fails
	if encoderror != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
	}
}
