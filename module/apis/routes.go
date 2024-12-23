package apis

import (
	"my-go-project/module/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// locationHandler := handlers.NewLocationHandler()
	weatherHandler := handlers.NewWeatherHandler()
	weatherHandler.RegisterRoutes(app)

}
