package global

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

func GetCountryFromStorage(alpha_2 string) (Country, bool) {
	// Check the map for the country
	val, ok := countrymap[alpha_2]
	// Return the country and the status code if the conutry is in the map
	return val, ok
}
