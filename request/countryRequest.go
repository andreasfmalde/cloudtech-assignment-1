package request

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
)

/*
Request a country based on three different search types;
name, cca2 and cca3. The spesified country will be
returned as a conutry struct
*/
func RequestCountryInfo(search string, searchType string) (global.Country, error) {
	// Checking to see if the country is already in storage
	if country, status := global.GetCountryFromStorage(search, searchType); status {
		// Returning the country struct from the storage
		return country, nil
	}
	// The country is not in storage, a request has to be made
	var url string
	// Decide upon which url to use based on search type
	if searchType == global.NAME_TYPE {
		url = global.COUNTRY_API_URL + "name/" + search + "?fullText=true"
	} else {
		url = global.COUNTRY_API_URL + "alpha/" + search
	}
	// Send the request
	res, err := SendGETRequest(url)
	// Handle error if request fails
	if err != nil {
		return global.Country{}, err
	}
	// Return a country struct based on the decoded JSON from the request
	return decodeCountryJSON(res)
}

/*
Return a country based on the JSON provided and decoded from
the http GET response in parameter
*/
func decodeCountryJSON(res *http.Response) (global.Country, error) {
	var countryList []global.Country
	// Decode JSON from the request into a country struct
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&countryList)
	// Handle error if decoding fails
	if err != nil {
		return global.Country{}, err
	}
	// Add conutry to storage to redure API calls
	global.AddCountryToStorage(countryList[0].CCA2, countryList[0])
	// Return the country
	return countryList[0], nil
}
