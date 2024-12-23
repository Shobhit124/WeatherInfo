package main

import (
	//fiber
	"my-go-project/module/apis"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main()	{
	//create server usiing fiber
	app := fiber.New()

	//setup routes
	apis.SetupRoutes(app)
	

	
	fmt.Println("Server running on :3000")
	//start server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
	
}
	
	


// type GeoResponse struct {
// 	Lat     float64 `json:"lat"`
// 	Lon     float64 `json:"lon"`
// 	Name    string  `json:"name"`
// 	State   string  `json:"state"`
// 	Country string  `json:"country"`
// }

// func main() {
// 	CityName := "Yamunanagar" // Example City
// 	StateCode := "HR"         // Example State Code
// 	CountryCode := "IN"       // Example Country Code

// 	// Construct API URL
// 	url := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&appid=d29386cf5d5075695804191b984fda64", CityName, StateCode, CountryCode)
// https://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={API_KEY}

// 	// Make the GET request
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal("Error while making the GET request:", err)
// 	}
// 	defer resp.Body.Close() // Ensure the response body is closed

// 	// Check if the request was successful
// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("Error: received non-OK HTTP status: %d", resp.StatusCode)

// 		// Read the response body
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatal("Error while reading the response body:", err)
// 		}

// 		// Decode the JSON response into a struct
// 		var geoResponse []GeoResponse
// 		err = json.Unmarshal(body, &geoResponse)
// 		if err != nil {
// 			log.Fatal("Error while unmarshalling JSON:", err)
// 		}

// 		// Print the parsed response
// 		fmt.Println("Parsed API Response:")
// 		for _, data := range geoResponse {
// 			encodedata, err := json.MarshalIndent(data, "", "  ")
// 			if err != nil {
// 				log.Fatal("Error while marshalling JSON:", err)
// 			}
// 			fmt.Println(string(encodedata))
// 			// fmt.Printf("City: %s, State: %s, Country: %s, Lat: %f, Lon: %f\n", data.Name, data.State, data.Country, data.Lat, data.Lon)
// 		}
// 	}
// }
