package routes

import "github.com/gofiber/fiber/v2"

func SendResponse(responseBody interface{}, httpStatus int, c *fiber.Ctx) error {
	return c.Status(httpStatus).JSON(responseBody)
}

type BaseResponseBody struct {
	ErrorMessage string `json:"errorMessage"`
	ResponseCode
}

type CreateAccountResponseBody struct {
	BaseResponseBody
	Email string
}
