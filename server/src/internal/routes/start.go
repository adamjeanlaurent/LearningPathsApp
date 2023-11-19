package routes

import (
	"encoding/json"
	"log"
	"os"

	"github.com/adamjeanlaurent/LearningPathsApp/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/jinzhu/gorm"
)

type RequestHandlerClient struct {
	DB *gorm.DB
}

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

func configureRoutes(app *fiber.App, handler *RequestHandlerClient) {
	var v1 fiber.Router = app.Group("v1/api")

	v1.Use(requestid.New())
	v1.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:   "${time}: ${locals:requestid} ${status} - ${method} ${path}\n",
		Output:   os.Stdout,
		TimeZone: "Local",
	}))
	v1.Use(logRequestBody)

	v1.Get("/createAccount", handler.CreateAccountRequestHandler)
}

func RunServer() {
	// connect to db
	var db *gorm.DB
	db, err := database.ConnectAndSetup()

	// check for db connection errors
	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("Nil database connection")
	}

	handler := &RequestHandlerClient{DB: db}

	var app *fiber.App = fiber.New()

	configureRoutes(app, handler)

	log.Fatal(app.Listen(":3000"))
}
