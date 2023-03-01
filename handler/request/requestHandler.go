package request

import (
	"assignment-1/global"
	"net/http"
)

func CombineUniversityAndCountry(uniList []global.University) ([]global.UniversityInformationStruct, error) {
	var universityAndCountryList []global.UniversityInformationStruct

	for _, uni := range uniList {
		country, err := RequestCountryInfo(uni.ISOcode, global.CCA2_TYPE)

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

func SendGETRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	defer client.CloseIdleConnections()

	return client.Do(req)

}
