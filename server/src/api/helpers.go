package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func sendParsingErrorResponse(responseBody IBaseResponseBody, c *fiber.Ctx) error {
	return sendResponse(responseBody, c, fiber.StatusBadRequest, ResponseCode_GenericError, "Parsing Failed")
}

func sendInternalServerError(responseBody IBaseResponseBody, c *fiber.Ctx, message string) error {
	return sendResponse(responseBody, c, fiber.StatusInternalServerError, ResponseCode_GenericError, message)
}

func sendSuccess(responseBody IBaseResponseBody, c *fiber.Ctx) error {
	return sendResponse(responseBody, c, fiber.StatusOK, ResponseCode_Success, "")
}

func sendResponse(responseBody IBaseResponseBody, c *fiber.Ctx, httpStatus int, responseCode ResponseCode, errorMessage string) error {
	responseBody.SetErrorMessage(errorMessage)
	responseBody.SetResponseCode(responseCode)
	return c.Status(httpStatus).JSON(responseBody)
}

func newBaseResponseBody() *BaseResponseBody {
	return &BaseResponseBody{
		ErrorMessage: "",
		ResponseCode: ResponseCode_Success,
	}
}

func parseRequestBody(out interface{}, c *fiber.Ctx) error {
	return c.BodyParser(&out)
}

func getUserStableIDFromContext(c *fiber.Ctx) (string, error) {
	// type assertion so we don't panic
	userStableId, ok := c.Locals("userStableId").(string)

	if !ok || userStableId == "" {
		return "", errors.New("invalid userStableId")
	}

	return userStableId, nil
}

func getUserTableIDFromContext(c *fiber.Ctx) (uint, error) {
	// type assertion so we don't panic
	userTableId, ok := c.Locals("userStableId").(uint)

	if !ok || userTableId == 0 {
		return 0, errors.New("invalid userTableId")
	}

	return userTableId, nil
}
