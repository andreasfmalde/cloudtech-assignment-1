package request

import (
	"assignment-1/global"
	"encoding/json"
)

func RequestUniversityInformation(url string) ([]global.UniversityInformationStruct, error) {

	res, err := SendGETRequest(url)

	if err != nil {
		return nil, err
	}

	var universityList []global.University

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&universityList)

	if err != nil {
		return nil, err
	}

	return CombineUniversityAndCountry(universityList)

}
