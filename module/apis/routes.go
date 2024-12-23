package apis

import (
	"my-go-project/module/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	locationHandler := handlers.NewLocationHandler()
	weatherHandler := handlers.NewWeatherHandler()
	appGroup := app.Group("/api/v1")
	appGroup.Post("/location", locationHandler.GetLocation)
	appGroup.Post("/weather", weatherHandler.GetWeather)
}