# Context-sensitive University Search Service
This project is the submission of Andreas Follevaag Malde in Cloud Technologies 2023. The project is a REST web application written in Golang where one can retrieve information about universities around the world. Some information of the country where the university is located will also be included.


The application uses two third-party APIs for retriving information of universities around the world and information of countries. The REST web services used are: 
- http://universities.hipolabs.com/
    - Documentation/Source can be found here: https://github.com/Hipo/university-domains-list/
- https://restcountries.com/
    - Documentation can be found here: https://gitlab.com/amatos/rest-countries

The hippolabs API are used primarily for retreving university information while the restcountries API are used to get more information of each country.

## Endpoints
This application has primarily three endpoints that can be used to retrieve information. These endpoint are: 
```
/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
``` 
### Default
---
When trying to reach another endpoint than these three, one will be met by a default message showing some information about the application, most importantly the URL for the available endpoints to use, as listed above. 

The default message looks like this:
```json
{
	"app-name": "Context-sensitive University Search Service",
	"madeby": "Andreas Follevaag Malde",
	"version": "v1.0",
	"endpoints": {
		"diagnostics": "/unisearcher/v1/diag/",
		"neighbourinfo": "/unisearcher/v1/neighbourunis/",
		"uniinfo": "/unisearcher/v1/uniinfo/"
	}
}
```

### University Information
---
The first endpoint will return information about one or many universities that has part of the search word in their names. An example is for the search word **norwegian**. The return output will include universities like "Norwegian State Academy of Music",  "Norwegian University of Science and Technology" etc. 
#### Request
The request to this endpoint should be in this format. 
```text
Method: GET
Path: unisearcher/v1/uniinfo/{:partial_or_complete_university_name}/
```
One has to include a university name in the ``` {:partial_or_complete_university_name} ```. Otherwise an error will be returned. An example of a valid request is:
``` text
/unisearcher/v1/uniinfo/University%20of%20Bergen
```

#### Response
The respose for the request above will look like this:
```json
[
	{
		"name": "University of Bergen",
		"country": "Norway",
		"isocode": "NO",
		"webpages": [
			"http://www.uib.no/"
		],
		"languages": {
			"nno": "Norwegian Nynorsk",
			"nob": "Norwegian Bokmål",
			"smi": "Sami"
		},
		"map": "https://www.openstreetmap.org/relation/2978650"
	}
]
```
The return status code will be: **200 OK**.
Other status codes that might be returned together with an error message from this endpoint are:
- **405** -  Method Not Allowed
- **403** -  Forbidden
- **500** -  Internal Server Error

### Neighbouring Countries Universities
---
The second endpoint will return an overview of universities of neighbouring countries to the country the search is based on, that has parts of their names matching the university name searched for. In this endpoint, the client also has the opportunity to limit how many universities to show, but that is optional.
#### Request
The request to the second endpoint should follow this format:
```text
Method: GET
Path: unisearcher/v1/neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}
```
- ```{:country_name}``` Mandatory, name of the country to base the search on
- ```{:partial_or_complete_university_name}``` Mandatory, university name to search for
- ```{?limit={:number}}``` Optional, set the limit of how many results to return

A valid request to this endpoint can look like this:
```
/unisearcher/v1/neighbourunis/norway/science?limit=2
```

#### Response
The response of the request above will look like this: 
```json
[
	{
		"name": "Häme University of Applied Sciences",
		"country": "Finland",
		"isocode": "FI",
		"webpages": [
			"https://www.hamk.fi/"
		],
		"languages": {
			"fin": "Finnish",
			"swe": "Swedish"
		},
		"map": "openstreetmap.org/relation/54224"
	},
	{
		"name": "Laurea University of Applied Sciences",
		"country": "Finland",
		"isocode": "FI",
		"webpages": [
			"http://www.laurea.fi/"
		],
		"languages": {
			"fin": "Finnish",
			"swe": "Swedish"
		},
		"map": "openstreetmap.org/relation/54224"
	}
]
```
The return status code will also here be **200 OK** if all went as it should.

Other status codes that could be returned if there are any error at this endpoint are:
- **405** -  Method Not Allowed
- **404** -  Not Found
- **500** -  Internal Server Error

### Diagnostics
If one is constantly experiencing errors in return on the two other endpoints, one can make a request to the diagnostics endpoint. This endpoint will return status codes for both third party APIs, as well as the version of the web application and the uptime of the service. 
#### Request 
The request to the diagnostics endpoint should follow this format:
```text
Method: GET
Path: unisearcher/v1/diag/
```
This endpoint takes no other input arguments than the path iteself. 
#### Response 
The response from the request above will look like this: 
```json
{
	"app-name": "Context-sensitive University Search Service",
	"madeby": "Andreas Follevaag Malde",
	"version": "v1.0",
	"endpoints": {
		"diagnostics": "/unisearcher/v1/diag/",
		"neighbourinfo": "/unisearcher/v1/neighbourunis/",
		"uniinfo": "/unisearcher/v1/uniinfo/"
	}
}
```
The return status code will also here be **200 OK** if everything went as it should. Other status code that can come if there are any errors at this endpoint are:
- **405** -  Method Not Allowed
- **500** -  Internal Server Error

## Deployment/Installation
---
There are two ways of running this application; spin up a server locally on your computer, or access the service from the web.
### Local Computer Installation
To run this application locally on your computer make sure Golang is installed on the system. Clone this repository and navigate into the cmd folder in the project. From there simply run the command
```bash
    go run main.go
```
### Deployment
The application is deployed to the web using the Render API for hosting. Because of this, the application is already up and runnig. It can be reached on this URL:

https://render.com/ 

## Special features
---
### Storage/Cache functionality
Using different third-party APIs to recieve different types of information may have some disadvantages. The amount of requests that has to be made to each of the APIs is one example. Let's says that one would like to search universities with "norwegian" in their name. Many of the univerities that are returned are in Norway, and for every university one would have to make a request to conutry API to get more information of the country the univeristy is located in. This will result in many requests to one of the APIs where the response will be the same country information every time. That is a problem.

The solution of this application is to make a storage to store information of each requested country. Before sending a request to the country API, the application will check the storage of the country already exists there. If it does, no request will be made. If it doesn't exist, a request will be made, and the returned result will then be placed in storage. This will lower the API calls drastically as one can see from this snippet from the log of the application where this was the url path ``` /unisearcher/v1/neighbourunis/norway/science?limit=5```.

Log:
```
2023/03/03 01:07:07 Must request country from API
2023/03/03 01:07:08 Must request country from API
2023/03/03 01:07:08 Must request country from API
2023/03/03 01:07:09 Must request country from API
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
2023/03/03 01:07:09 Country found in storage
...
```
Only 4 requests had to be made to the country API, instead of **24** requests originally. Sending the same request again, will lead to 0 country API calls because all the countries are in storage. The storage will not hold more than around 200 countries, so most modern computer will have no problem storing them, even though it is stored in memory. 

## Technologies
---
- Golang version 1.18
- Render API - for hosting
