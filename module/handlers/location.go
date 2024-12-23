package handlers

import (
	"fmt"
	"my-go-project/module/dtos"
	"my-go-project/module/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LocationHandler struct {
	Validator *validator.Validate
}

func NewLocationHandler() *LocationHandler {
	return &LocationHandler{
		Validator: validator.New(), // Initialize the validator
	}
}

func (h *LocationHandler) GetLocation(ctx *fiber.Ctx) error {
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
	responseData, err := h.fetchLocation(ctx, req)

	if err != nil {
		// Return a 500 Internal Server Error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ApiStandardResponseModel{
			Success: false,
			Error:   err.Error(),
		})
	}

	// Return a success response
	return ctx.JSON(dtos.ApiStandardResponseModel{
		Success: true,
		Data:    responseData,
	})
}

func (h *LocationHandler) fetchLocation(ctx *fiber.Ctx, req dtos.LocationRequest) (*dtos.LocationResponse, error) {
	// Add your logic here, e.g., save the location, process data, etc.
	responseData, err := services.GetLocation(req)

	if err != nil {
		// Return a 500 Internal Server Error response
		return nil, ctx.Status(fiber.StatusInternalServerError).JSON(dtos.ApiStandardResponseModel{
			Success: false,
			Error:   err.Error(),
		})
	}

	return responseData, nil

}
