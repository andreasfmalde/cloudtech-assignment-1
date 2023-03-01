package global

// Resource root paths
const DEFAULT_PATH = "/"
const UNI_INFO_PATH = "/unisearcher/v1/uniinfo/"
const NEIGHBOUR_UNI_PATH = "/unisearcher/v1/neighbourunis/"
const DIAG_PATH = "/unisearcher/v1/diag/"

// Time variable - NOT CONSTANT
// Used to store the start time of when the service started
var START_TIME int64

// Default port
const DEFAULT_PORT = "8080"

// API URLs
const UNIVERSITY_API_URL = "http://universities.hipolabs.com/"
const COUNTRY_API_URL = "https://restcountries.com/v3.1/"
