package dtos

type LocationRequest struct {
	City        string `json:"city" validate:"required"`
	StateCode   string `json:"stateCode" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}
