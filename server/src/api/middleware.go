package api

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"

	"github.com/adamjeanlaurent/LearningPathsApp/utility"
	"github.com/gofiber/fiber/v2"
)

func validateRequestBody(expectedType interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Create an instance of the expected type
		requestBody := reflect.New(reflect.TypeOf(expectedType).Elem()).Interface()

		// Parse the JSON request body into the expected type
		if err := ctx.BodyParser(requestBody); err != nil {
			utility.LogError(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Parsing Error",
			})
		}

		// You can now access the parsed data as requestBody
		ctx.Locals("body", requestBody)

		// Continue processing the request
		return ctx.Next()
	}
}

func (server *ApiServer) validateJwtToken(c *fiber.Ctx) error {
	var jwtToken = c.Cookies("jwt")
	stableId, userTableId, err := utility.ParseJwt(jwtToken, server.config.GetJwtSecretKey())
	if err != nil {
		return err
	}

	c.Locals("userStableId", stableId)
	c.Locals("userTableID", userTableId)

	stableId, ok := c.Locals("userStableId").(string)

	if !ok || stableId == "" {
		utility.LogError(errors.New("invalid user info in auth token"))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	tableId, ok := c.Locals("userTableId").(uint)

	if !ok || tableId == 0 {
		utility.LogError(errors.New("invalid user info in auth token"))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return nil
}

func (server *ApiServer) logRequestBody(c *fiber.Ctx) error {
	body := make(map[string]interface{})

	// Read the request body
	var bodyBytes []byte = c.Body()

	if len(bodyBytes) == 0 {
		log.Printf("empty json body")
		return c.Next()
	}

	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return err
	}

	// Convert the map to JSON for logging
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error converting body to JSON: %v", err)
		return err
	}

	// Log the JSON body
	log.Printf("Request Body: %s", bodyJSON)

	return c.Next()
}
