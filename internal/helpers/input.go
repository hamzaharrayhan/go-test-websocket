package helpers

import "github.com/go-playground/validator/v10"

type PhoneNumberInput struct {
	Number     string `form:"number" binding:"required"`
	ProviderID string `form:"provider_id" binding:"required"`
}

type PhoneNumberResponse struct {
	ID         string `json:"id"`
	Number     string `json:"number"`
	ProviderID string `json:"provider_id"`
}

type ProviderResponse struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func FormatValidationError(err error) []string {
	var errors []string

	// loop errors
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func JSONResponse(status string, data interface{}) Response {
	jsonResponse := Response{
		Status: status,
		Data:   data,
	}

	return jsonResponse
}
