package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
)

/*
Default path handler will just display an error message for the
user, telling the user that the site does not exist
*/
func DefaultPathHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure only GET methods are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "This site only allows "+http.MethodGet+" methods.", http.StatusMethodNotAllowed)
		return
	}
	// Add content-type json to the response header
	w.Header().Add("content-type", "application/json")

	// Create an endpoint map to be used in the default endpoint message
	endpoints := make(map[string]string)
	endpoints["uniinfo"] = "/unisearcher/v1/uniinfo/"
	endpoints["neighbourinfo"] = "/unisearcher/v1/neighbourunis/"
	endpoints["diagnostics"] = "/unisearcher/v1/diag/"
	// Create the default endpoint message
	message := global.DefaultMessage{
		Name:     "Context-sensitive University Search Service",
		MadeBy:   "Andreas Follevaag Malde",
		Version:  "v1.0",
		Endpoint: endpoints,
	}
	// Send the default message as JSON to the client
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(message)

	// Display an error to the user if the encoding fails
	if err != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
	}
}
