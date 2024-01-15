package api

import (
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

func getUserStableIDFromContext(c *fiber.Ctx) string {
	userStableId := c.Locals("userStableId").(string)

	return userStableId
}

func getUserTableIDFromContext(c *fiber.Ctx) uint {
	userTableId := c.Locals("userStableId").(uint)

	return userTableId
}
