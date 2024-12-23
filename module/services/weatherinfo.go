package services

import (
	"encoding/json"
	"fmt"
	"io"
	"my-go-project/module/dtos"
	"net/http"
)


func GetWeather(request dtos.LocationResponse) (*dtos.WeatherResponse, error) {
	// Construct the API URL
	api := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=d29386cf5d5075695804191b984fda64", request.Latitude, request.Longitude)

	// Make the GET request
	resp, err := http.Get(api)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close() // Ensure response body is closed

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response JSON into WeatherResponse struct
	var weatherData dtos.WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	return &weatherData, nil
}

