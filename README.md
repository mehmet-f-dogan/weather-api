# Weather API

A simple API built in Go that retrieves the current temperature for a given latitude and longitude.

## Installation

- Clone this repository to your local machine.
- Install Go on your machine.

## Usage

1. To start the server, run the following command in the project directory:
<br>`go run main.go`<br>
2. The server will start on port 8080 by default. To specify a different port, set the PORT environment variable:
<br>`PORT=8888 go run main.go`<br>
3. The API can be accessed at [localhost:8080](http://localhost:8080/api). The latitude and longitude can be specified as query parameters:
<br>http://localhost:8080/api?lat=52.52&lon=13.41<br>
4. The API will return a JSON response with the current temperature:
<br>`{ "temp": 3.4 }`<br>  
5. If an error occurs, the API will return a JSON error response with a 400 status code:
<br>`{
"error": true,
"reason": "Missing lat/lon parameters"
}`

## Testing

To run the unit tests for the program, run the following command in the project directory:  
`go test`  
This will run the tests defined in main_test.go and output the test results.
