package controller

import (
	"jti-test/internal/helpers"
	"jti-test/internal/services"

	"github.com/gofiber/fiber/v2"
)

var phoneNumberService services.PhoneNumberService
var providerService services.ProviderService

func NewPhoneNumberController(phoneNumber services.PhoneNumberService, provider services.ProviderService) {
	phoneNumberService = phoneNumber
	providerService = provider
}

func AddNewPhoneNumber(c *fiber.Ctx) error {
	var input helpers.PhoneNumberInput

	if err := c.BodyParser(&input); err != nil {
		response := helpers.JSONResponse("failed", err)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// send to service
	newPhoneNumber, err := phoneNumberService.CreatePhoneNumber(input)

	if err != nil {
		response := helpers.JSONResponse("failed", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	newPhoneNumberResponse := helpers.PhoneNumberResponse{
		ID:         newPhoneNumber.ID.String(),
		Number:     newPhoneNumber.Number,
		ProviderID: newPhoneNumber.ProviderID.String(),
	}

	response := helpers.JSONResponse("created", newPhoneNumberResponse)
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAllProvider(c *fiber.Ctx) error {
	res, err := providerService.GetAll()
	if err != nil {
		response := helpers.JSONResponse("failed", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helpers.JSONResponse("retrieved", res)
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAllPhoneNumber(c *fiber.Ctx) error {
	res, err := phoneNumberService.GetAll()
	if err != nil {
		response := helpers.JSONResponse("failed", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helpers.JSONResponse("retrieved", res)
	return c.Status(fiber.StatusOK).JSON(response)
}

func UpdatePhoneNumber(c *fiber.Ctx) error {
	var input helpers.PhoneNumberInput

	id := c.Params("id", "0")
	if err := c.BodyParser(&input); err != nil {
		response := helpers.JSONResponse("failed", err)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// send to service
	newPhoneNumber, err := phoneNumberService.Update(id, input)

	if err != nil {
		errorMessages := helpers.FormatValidationError(err)

		response := helpers.JSONResponse("failed", errorMessages)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	newPhoneNumberResponse := helpers.PhoneNumberResponse{
		ID:         id,
		Number:     newPhoneNumber.Number,
		ProviderID: newPhoneNumber.ProviderID.String(),
	}

	response := helpers.JSONResponse("updated", newPhoneNumberResponse)
	return c.Status(fiber.StatusOK).JSON(response)
}

func DeletePhoneNumber(c *fiber.Ctx) error {
	id := c.Params("id", "0")
	_, err := phoneNumberService.Delete(id)
	if err != nil {
		response := helpers.JSONResponse("failed", err)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	res := helpers.JSONResponse("deleted", nil)

	return c.Status(fiber.StatusOK).JSON(res)
}
