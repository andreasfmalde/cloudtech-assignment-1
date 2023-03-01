package global

import (
	"strings"
)

// Storage collection to store all country structs to reduce api calls
var countrymap = map[string]Country{}

/*
* Add a country struct to the sotrage collection to reduce
* api calls. If country already is in the collection, false is returned
 */
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

/*
* Return a country struct from the storage collection based on
* search word and search type. False is returned if country is
* not in collecion
 */
func GetCountryFromStorage(search string, searchType string) (Country, bool) {
	switch searchType {
	case "cca2":
		// Check the map for the country
		val, ok := countrymap[search]
		// Return the country and the status code if the conutry is in the map
		return val, ok
	case "cca3":
		// Loop through all countries in the map
		for _, country := range countrymap {
			//Return conutry if CCA3 is a match
			if country.CCA3 == search {
				return country, true
			}
		}
	case "name":
		// Loop through all countries in the map
		for _, country := range countrymap {
			name := country.Name["common"].(string) // Retrieve the common name of the country
			//Return conutry if name is a match
			if strings.ToLower(name) == strings.ToLower(search) {
				return country, true
			}
		}
	}
	return Country{}, false // Return false if country is not in the storage
}
