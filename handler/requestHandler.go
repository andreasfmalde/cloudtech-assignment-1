package handler

import (
	"assignment-1/global"
	"encoding/json"
	"net/http"
	"strings"
)

func RequestUniversityInformation(search string) ([]global.University, error) {
	search = strings.Replace(search, " ", "%20", -1)

	url := global.UNIVERSITY_API_URL + "search?name=" + search

	r, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	var universityList []global.University

	decoder := json.NewDecoder(res.Body)

	er := decoder.Decode(&universityList)

	if er != nil {
		return nil, er
	}

	return universityList, nil

}
