package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
)

func RequestUniversityInformation(url string) ([]global.UniversityInformationStruct, error) {

	res, err := sendGetRequest(url)

	if err != nil {
		return nil, err
	}

	var universityList []global.University

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&universityList)

	if err != nil {
		return nil, err
	}

	return combineUniversityAndCountry(universityList)

}

func combineUniversityAndCountry(uniList []global.University) ([]global.UniversityInformationStruct, error) {
	var universityAndCountryList []global.UniversityInformationStruct

	for _, uni := range uniList {
		country, err := requestCountryInfo(uni.ISOcode, global.CCA2_TYPE)

		if err != nil {
			return nil, err
		}

		uniAndCountry := global.UniversityInformationStruct{Name: uni.Name,
			Country: uni.Country, ISOcode: uni.ISOcode,
			Webpages: uni.Webpages, Languages: country.Languages,
			Map: country.Map["openStreetMaps"]}

		universityAndCountryList = append(universityAndCountryList, uniAndCountry)
	}

	return universityAndCountryList, nil

}

func sendGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	defer client.CloseIdleConnections()

	return client.Do(req)

}

func requestCountryInfo(search string, searchType string) (global.Country, error) {
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

	res, err := sendGetRequest(url)
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
