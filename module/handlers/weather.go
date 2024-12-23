package handlers

import (
	"fmt"
	"my-go-project/module/dtos"
	"my-go-project/module/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type WeatherHandler struct {
	Validator       *validator.Validate
	LocationHandler *LocationHandler
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{
		Validator: validator.New(), // Initialize the validator
	}
}

func (h *WeatherHandler) RegisterRoutes(app *fiber.App)  {
	appGroup := app.Group("/api/v1")
	// appGroup.Post("/location", locationHandler.GetLocation)
	appGroup.Post("/weather", h.GetWeather)
	appGroupV2 := app.Group("/api/v2")
	appGroupV2.Post("/location", h.LocationHandler.GetLocation)
}

func (h *WeatherHandler) GetWeather(ctx *fiber.Ctx) error {

	var req dtos.LocationRequest

	// Parse the request body into the struct
	if err := ctx.BodyParser(&req); err != nil {
		// Log the error for debugging purposes
		fmt.Printf("Error parsing body: %v\n", err)

		// Return a 400 Bad Request response
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.ApiStandardResponseModel{
			Success: false,
			Error:   "Failed to parse request body",
		})
	}

	// Validate the parsed data
	if err := h.Validator.Struct(&req); err != nil {
		// Extract detailed validation errors
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))

		for i, fieldErr := range validationErrors {
			errorMessages[i] = fmt.Sprintf("Field '%s' failed validation on '%s' rule", fieldErr.Field(), fieldErr.Tag())
		}

		// Return a 422 Unprocessable Entity response
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(dtos.ApiStandardResponseModel{
			Success: false,
			Error:   "Validation failed",
			Data:    errorMessages,
		})
	}

	// Add your logic here, e.g., save the location, process data, etc.
	responseData, _ := h.LocationHandler.fetchLocation(ctx,req)


	weatherService, err := services.GetWeather(*responseData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ApiStandardResponseModel{
			Success: false,
			Error:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dtos.ApiStandardResponseModel{
		Success: true,
		Data:    weatherService,
	})

}