package global

/*
	Struct used to store and represent university information
*/
type UniversityInformationStruct struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	ISOcode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
}

type University struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	ISOcode  string   `json:"alpha_two_code"`
	Webpages []string `json:"web_pages"`
}

type Country struct {
	Languages map[string]string `json:"languages"`
	Borders   []string          `json:"borders"`
	Map       map[string]string `json:"maps"`
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
