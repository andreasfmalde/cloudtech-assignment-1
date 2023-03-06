package global

/*
Some country names are not the same in the two APIs. This map contains
the convertion of country names that are not the same.
The keys of the map are country names from the REST Countries API,
while the values are the HipoLabs API name of the country
*/
var countryNames = map[string]string{
	"Bolivia":                "Bolivia, Plurinational State of",
	"British Virgin Islands": "Virgin Islands, British",
	"Brunei":                 "Brunei Darussalam",
	"Czechia":                "Czech Republic",
	"DR Congo":               "Congo, the Democratic Republic of the",
	"Eswatini":               "Swaziland",
	"Ivory Coast":            "Côte d'Ivoire",
	"Laos":                   "Lao People's Democratic Republic",
	"Macau":                  "Macao",
	"Moldova":                "Moldova, Republic of",
	"North Korea":            "Korea, Democratic People's Republic of",
	"Palestine":              "Palestine, State of",
	"Republic of the Congo":  "Congo",
	"Russia":                 "Russian Federation",
	"South Korea":            "Korea, Republic of",
	"Syria":                  "Syrian Arab Republic",
	"Tanzania":               "Tanzania, United Republic of",
	"Vatican City":           "Holy See (Vatican City State)",
	"Venezuela":              "Venezuela, Bolivarian Republic of",
	"Vietnam":                "Viet Nam",
}

/*
A function used to convert country names from the REST Conutries
API to valid HipoLabs API country names. If country is not in the
map, false will be returned as the second return argument
*/
func ConvertCountryName(country string) (string, bool) {
	val, ok := countryNames[country]
	return val, ok
}
