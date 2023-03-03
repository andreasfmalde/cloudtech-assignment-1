package request

import (
	"assignment-1/global"
	"net/http"
)

/*
Retrieve country information of each university in the university list
provided and combine the country information and the university
into a new struct
*/
func CombineUniversityAndCountry(uniList []global.University) ([]global.UniversityInformationStruct, error) {
	var universityAndCountryList []global.UniversityInformationStruct
	// Loop throuh all universities in the uniList
	for _, uni := range uniList {
		country, err := RequestCountryInfo(uni.ISOcode, global.CCA2_TYPE)
		// Terminate function if a fail in retrieveing country info occurs
		if err != nil {
			return nil, err
		}
		// Combine the university and the country struct into a new struct
		uniAndCountry := global.UniversityInformationStruct{Name: uni.Name,
			Country: uni.Country, ISOcode: uni.ISOcode,
			Webpages: uni.Webpages, Languages: country.Languages,
			Map: country.Map["openStreetMaps"]}
		// Append the struct to the common list
		universityAndCountryList = append(universityAndCountryList, uniAndCountry)
	}
	// Return the combined university and country list
	return universityAndCountryList, nil

}

/*
Send a get request to a specified URL and return
the respone
*/
func SendGETRequest(url string) (*http.Response, error) {
	// Create a new http request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// Terminate the function if a fail occurs
	if err != nil {
		return nil, err
	}
	// Create a client and defer the connection
	client := &http.Client{}
	defer client.CloseIdleConnections()
	// Send the request and return the response
	return client.Do(req)
}
