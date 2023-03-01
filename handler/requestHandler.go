package handler

import (
	"assignment-1/global"
	"encoding/json"
	"log"
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
		country, err := requestCountryInfoByAlpha(uni.ISOcode, "cca2")

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

func requestCountryInfoByAlpha(alpha_2 string, alphaType string) (global.Country, error) {
	// Checking to see if the country is already in storage
	if country, status := global.GetCountryFromStorage(alpha_2, alphaType); status {
		// Returning the country struct from the storage
		log.Println("Country found in storage")
		return country, nil
	}
	// The country is not in storage, a request has to be made
	url := global.COUNTRY_API_URL + "alpha/" + alpha_2
	log.Println("----REquest---------")
	res, err := sendGetRequest(url)
	// Handle error if request failed
	if err != nil {
		return global.Country{}, err
	}

	var countryList []global.Country

	decoder := json.NewDecoder(res.Body)

	err1 := decoder.Decode(&countryList)

	if err1 != nil {
		return global.Country{}, err1
	}
	log.Println(countryList[0].Name["common"])
	// Add the country to storage to reduce api calls to the same country
	if alphaType == "cca2" {
		global.AddCountryToStorage(alpha_2, countryList[0])
	} else {
		global.AddCountryToStorage(countryList[0].CCA2, countryList[0])
	}

	return countryList[0], nil
}

func requestCountryInfoByName(name string) (global.Country, error) {
	if country, status := global.GetCountryFromStorage(name, "name"); status {
		// Returning the country struct from the storage
		log.Println("Country found in storage")
		return country, nil
	}

	url := global.COUNTRY_API_URL + "name/" + name + "?fullText=true"

	res, err := sendGetRequest(url)

	if err != nil {
		return global.Country{}, err
	}

	var countryList []global.Country

	decoder := json.NewDecoder(res.Body)

	err1 := decoder.Decode(&countryList)

	if err1 != nil {
		return global.Country{}, err1
	}
	return countryList[0], nil
}
