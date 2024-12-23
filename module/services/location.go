package services

import (
	"encoding/json"
	"fmt"
	"io"
	"my-go-project/module/dtos"
	"net/http"
)


func GetLocation(request dtos.LocationRequest) (*dtos.LocationResponse, error) {
	// Construct the API URL
	api := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&appid=d29386cf5d5075695804191b984fda64", request.City, request.StateCode, request.CountryCode)

	// Make the GET request
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Ensure response body is closed

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response JSON into LocationResponse struct
	var locData []dtos.LocationResponse
	if err := json.Unmarshal(body, &locData); err != nil {
		return nil, err
	}

	// If no data is found, return an error
	if len(locData) == 0 {
		return nil, fmt.Errorf("no location foud")
	}
	fmt.Println(locData[0])
	// Return the success response with the parsed data (only the first result, as it's a list)
	return   &locData[0], nil
}
