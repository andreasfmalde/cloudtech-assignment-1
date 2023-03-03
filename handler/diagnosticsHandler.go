package handler

import (
	"assignment-1/global"
	"assignment-1/handler/request"
	"encoding/json"
	"net/http"
	"time"
)

/*
The diagnostics handler provides all operations of the diag/
endpoint. This includes getting the status of the third-party
apis to see if they are available, as well as version and
uptime of the server
*/
func DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure only GET methods are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "This site only allows "+http.MethodGet+" methods.", http.StatusMethodNotAllowed)
		return
	}
	// Add content-type json to the response header
	w.Header().Add("content-type", "application/json")

	// Send a GET request to the University API
	universityAPI, err := request.SendGETRequest(global.UNIVERSITY_API_URL)
	if err != nil {
		http.Error(w, "Something went wrong in the universityAPI request", http.StatusInternalServerError)
		return
	}
	// Send a GET request to the Country API
	countryAPI, err1 := request.SendGETRequest(global.UNIVERSITY_API_URL)
	if err1 != nil {
		http.Error(w, "Something went wrong in the countryAPI request", http.StatusInternalServerError)
		return
	}
	// Make a diagnostics struct based on api statuses, version and uptime of server
	diagnostics := global.DiagnostictStruct{
		UniversityAPI: universityAPI.Status,
		CountryAPI:    countryAPI.Status,
		Version:       "v1",
		Uptime:        int(time.Now().Unix() - global.START_TIME),
	}

	// Encode the struct to JSON and send to client
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t") // Make JSON prettier to read
	err2 := encoder.Encode(diagnostics)
	// Handle any errors that may occur while encodig the JSON
	if err2 != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
		return
	}

}
