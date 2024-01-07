package api

import (
	"log"
	"os"

	"github.com/adamjeanlaurent/LearningPathsApp/storage"
	"github.com/adamjeanlaurent/LearningPathsApp/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func configureRoutes(server *ApiServer) {
	var v1 fiber.Router = server.app.Group("v1/api")

	v1.Use(requestid.New())
	v1.Use(logger.New(logger.Config{
		Format:   "${time}: ${locals:requestid} ${status} - ${method} ${path}\n",
		Output:   os.Stdout,
		TimeZone: "Local",
	}))
	v1.Use(server.logRequestBody)

	var authRouter fiber.Router = v1.Group("/auth")
	authRouter.Post("/createAccount", server.handleCreateAccount)
	authRouter.Get("/loginToAccount", server.handleLogin)

	var learningPathRouter fiber.Router = v1.Group("/learningPath")
	learningPathRouter.Use(server.validateJwtToken)
	learningPathRouter.Post("/create", server.handleCreateLearningPath)
	learningPathRouter.Post("/update/title", server.handleSetLearningPathTitle)

	var learningPathStopRouter fiber.Router = v1.Group("/learningPathStop")
	learningPathStopRouter.Use(server.validateJwtToken)
	learningPathStopRouter.Post("/create", server.handleCreateLearningPathStop)
	learningPathStopRouter.Post("/update/title", server.handleSetLearningPathStopTitle)
	learningPathStopRouter.Post("/update/body", server.handleSetLearningPathStopBody)
}

func (server *ApiServer) ConnectAndRun() {
	// connect to the db
	server.store = storage.NewMySqlStore()

	var err error = server.store.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// load config
	server.config = utility.NewServerConfiguration()

	err = server.config.Validate()

	if err != nil {
		log.Fatal(err)
	}

	server.app = fiber.New()

	configureRoutes(server)

	server.port = ":3000"

	log.Fatal(server.app.Listen(server.port))
}
