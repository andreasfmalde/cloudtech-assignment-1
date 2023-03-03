package request

import (
	"assignment-1/global"
	"encoding/json"
)

/*
Based on the URL, a list of universities with a spesified name
and optionally from a spesified country is returned
*/
func RequestUniversityInformation(url string) ([]global.UniversityInformationStruct, error) {
	// Send a GET request to the spesified API
	res, err := SendGETRequest(url)
	// Terminate the function if a fail in the GET request occurs
	if err != nil {
		return nil, err
	}

	var universityList []global.University
	// Decode JSON into a list of universities
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&universityList)
	// Terminate the function if the JSON decoding fails
	if err != nil {
		return nil, err
	}
	// Call another function to combine univeristy and country information
	return CombineUniversityAndCountry(universityList)

}
