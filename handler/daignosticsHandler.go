package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
	"time"
)

/*Some action*/
func DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "This site only allows "+http.MethodGet+" methods.", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("content-type", "application/json")

	universityAPI, err := sendGetRequest(global.UNIVERSITY_API_URL)
	if err != nil {
		http.Error(w, "Something went in the universityAPI request", http.StatusInternalServerError)
		return
	}
	countryAPI, err1 := sendGetRequest(global.UNIVERSITY_API_URL)
	if err1 != nil {
		http.Error(w, "Something went in the universityAPI request", http.StatusInternalServerError)
		return
	}

	diagnostics := global.DiagnostictStruct{
		UniversityAPI: universityAPI.Status,
		CountryAPI:    countryAPI.Status,
		Version:       "v1",
		Uptime:        int(time.Now().Unix() - global.START_TIME),
	}

	encoder := json.NewEncoder(w)

	err2 := encoder.Encode(diagnostics)

	if err2 != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
		return
	}

}
