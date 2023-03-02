package global

/*
	Struct used to store and represent university and country information
*/
type UniversityInformationStruct struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	ISOcode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
}

/*
	Struct used to store and represent university information
*/
type University struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	ISOcode  string   `json:"alpha_two_code"`
	Webpages []string `json:"web_pages"`
}

/*
	Struct used to store and represent country information
*/
type Country struct {
	Name      map[string]interface{} `json:"name"`
	CCA2      string                 `json:"cca2"`
	CCA3      string                 `json:"cca3"`
	Languages map[string]string      `json:"languages"`
	Borders   []string               `json:"borders"`
	Map       map[string]string      `json:"maps"`
}

/*
	Struct used to store and represent diagnostics information
*/
type DiagnostictStruct struct {
	UniversityAPI string `json:"universitiesapi"`
	CountryAPI    string `json:"countriesapi"`
	Version       string `json:"version"`
	Uptime        int    `json:"uptime"`
}

/*
	Struct used to display application info to the default endpoint
*/
type DefaultMessage struct {
	Name     string            `json:"app-name"`
	MadeBy   string            `json:"madeby"`
	Version  string            `json:"version"`
	Endpoint map[string]string `json:"endpoints"`
}
