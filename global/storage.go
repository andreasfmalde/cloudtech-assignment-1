package global

import "strings"

// Storage collection to store all country structs to reduce api calls
var countrymap = map[string]Country{}

func AddCountryToStorage(alpha_2 string, country Country) bool {
	// Check if country is already in the map
	if _, ok := countrymap[alpha_2]; ok {
		// The country is already in the map
		return false
	} else {
		//The country is not already in the map, it will be added
		countrymap[alpha_2] = country
		return true
	}
}

func GetCountryFromStorage(search string, searchType string) (Country, bool) {
	switch searchType {
	case "cca2":
		// Check the map for the country
		val, ok := countrymap[search]
		// Return the country and the status code if the conutry is in the map
		return val, ok
	case "cca3":
		for _, country := range countrymap {
			if country.CCA3 == search {
				return country, true
			}
		}
	case "name":
		for _, country := range countrymap {
			name := country.Name["common"].(string)
			if strings.ToLower(name) == strings.ToLower(search) {
				return country, true
			}
		}
	}
	return Country{}, false
}
