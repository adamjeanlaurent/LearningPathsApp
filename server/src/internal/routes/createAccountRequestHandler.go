package routes

import "github.com/gofiber/fiber/v2"

func (handler *RequestHandlerClient) CreateAccountRequestHandler(c *fiber.Ctx) error {
	return c.SendString("hi")
}
