package request

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
)

func RequestCountryInfo(search string, searchType string) (global.Country, error) {
	// Checking to see if the country is already in storage
	if country, status := global.GetCountryFromStorage(search, searchType); status {
		// Returning the country struct from the storage
		return country, nil
	}
	// The country is not in storage, a request has to be made
	var url string
	if searchType == global.NAME_TYPE {
		url = global.COUNTRY_API_URL + "name/" + search + "?fullText=true"
	} else {
		url = global.COUNTRY_API_URL + "alpha/" + search
	}

	res, err := SendGETRequest(url)
	// Handle error if request failed
	if err != nil {
		return global.Country{}, err
	}

	return decodeCountryJSON(res)
}

func decodeCountryJSON(res *http.Response) (global.Country, error) {
	var countryList []global.Country

	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&countryList)

	if err != nil {
		return global.Country{}, err
	}
	global.AddCountryToStorage(countryList[0].CCA2, countryList[0])
	return countryList[0], nil
}
