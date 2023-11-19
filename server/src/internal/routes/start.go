package routes

import (
	"log"

	"github.com/adamjeanlaurent/LearningPathsApp/internal/database"
	"github.com/gofiber/fiber/v2"

	"github.com/jinzhu/gorm"
)

type RequestHandlerClient struct {
	DB *gorm.DB
}

func configureRoutes(app *fiber.App, handler *RequestHandlerClient) {
	var v1 fiber.Router = app.Group("v1/api")
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
