package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
	"strings"
)

func RequestUniversityInformation(search string) ([]global.UniversityInformationStruct, error) {
	search = strings.Replace(search, " ", "%20", -1)

	url := global.UNIVERSITY_API_URL + "search?name=" + search

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
		country, err := requestCountryInfoByAlpha2(uni.ISOcode)

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

func requestCountryInfoByAlpha2(alpha_2 string) (global.Country, error) {
	url := global.COUNTRY_API_URL + "alpha/" + alpha_2

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
