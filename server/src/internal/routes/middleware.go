package routes

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// func authenicate(c *fiber.Ctx) error {

// }

func logRequestBody(c *fiber.Ctx) error {
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
