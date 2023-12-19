package api

import (
	"github.com/adamjeanlaurent/LearningPathsApp/storage"
	"github.com/adamjeanlaurent/LearningPathsApp/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Server interface {
	ConnectAndRun()
	handleCreateAccount(c *fiber.Ctx) error
	handleLogin(c *fiber.Ctx) error
	handleCreateLearningPath(c *fiber.Ctx) error
}

type ApiServer struct {
	port   string
	store  storage.Store
	config utility.Configuration
	app    *fiber.App
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		port:   "",
		store:  nil,
		config: nil,
		app:    nil,
	}
}

func (server *ApiServer) handleCreateAccount(c *fiber.Ctx) error {
	var requestBody CreateAccountRequestBody

	var response *CreateAccountResponseBody = NewCreateAccountResponseBody()

	var err error = parseRequestBody(&requestBody, c)
	if err != nil || requestBody.Email == "" || requestBody.Password == "" {
		utility.LogError(err)
		return sendParsingErrorResponse(response, c)
	}

	existingUser := storage.User{}

	queryResult := server.store.GetUserByEmail(requestBody.Email, &existingUser)

	if queryResult.Error != nil && queryResult.Error != gorm.ErrRecordNotFound {
		utility.LogError(queryResult.Error)
		return sendResponse(response, c, fiber.StatusOK, ResponseCode_AccountWithEmailAlreadyExists, "")
	}

	if existingUser.Email != "" {
		response.ResponseCode = ResponseCode_AccountWithEmailAlreadyExists
		return sendResponse(response, c, fiber.StatusOK, ResponseCode_AccountWithEmailAlreadyExists, "")
	}

	// TODO: check for valid email

	passwordHash, err := utility.HashPassword(requestBody.Password)

	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Password hashing failed")
	}

	creationResult, user := server.store.CreateUser(requestBody.Email, passwordHash)

	if creationResult.Error != nil {
		utility.LogError(creationResult.Error)
		return sendInternalServerError(response, c, "Failed to save user in DB")
	}

	jwtToken, err := utility.CreateNewJwt(user.StableId, user.ID, server.config.GetJwtSecretKey())

	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Failed to generate Jwt")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	return sendSuccess(response, c)
}

func (server *ApiServer) handleCreateLearningPath(c *fiber.Ctx) error {
	var requestBody CreateLearningPathRequestBody

	response := NewCreateLearningPathResponseBody()

	var err error = parseRequestBody(&requestBody, c)
	if err != nil || requestBody.Title == "" {
		utility.LogError(err)
		return sendParsingErrorResponse(response, c)
	}

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "invalid table ID")
	}

	creationResult, _ := server.store.CreateLearningPath(requestBody.Title, userTableID)

	if creationResult.Error != nil {
		utility.LogError(creationResult.Error)
		return sendInternalServerError(response, c, "Failed to save learning path in DB")
	}

	return sendSuccess(response, c)
}

func (server *ApiServer) handleLogin(c *fiber.Ctx) error {
	var requestBody LoginToAccountRequestBody

	var response *LoginToAccountResponseBody = NewLoginToAccountResponseBody()

	var err error = parseRequestBody(&requestBody, c)
	if err != nil || requestBody.Email == "" || requestBody.Password == "" {
		utility.LogError(err)
		return sendParsingErrorResponse(response, c)
	}

	existingUser := &storage.User{}

	queryResult := server.store.GetUserByEmail(requestBody.Email, existingUser)

	if queryResult.Error != nil {
		utility.LogError(queryResult.Error)
		return sendResponse(response, c, fiber.StatusUnauthorized, ResponseCode_NoAccountWithEmailFound, "")
	}

	err = utility.ComparePasswordWithHash(requestBody.Password, existingUser.Hash)
	if err != nil {
		utility.LogError(err)
		return sendResponse(response, c, fiber.StatusUnauthorized, ResponseCode_GenericError, "")
	}

	jwtToken, err := utility.CreateNewJwt(existingUser.StableId, existingUser.ID, server.config.GetJwtSecretKey())
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Failed to generate Jwt")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	return sendSuccess(response, c)
}
