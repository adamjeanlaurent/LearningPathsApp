package api

import (
	"github.com/adamjeanlaurent/LearningPathsApp/storage"
	"github.com/adamjeanlaurent/LearningPathsApp/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type ApiServer struct {
	port   string
	store  *storage.MySqlStore
	config *utility.ServerConfiguration
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
	requestBody := c.Locals("body").(*CreateAccountRequestBody)

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

	user, err := server.store.CreateUser(requestBody.Email, passwordHash)

	if err != nil {
		utility.LogError(err)
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

func (server *ApiServer) handleLogin(c *fiber.Ctx) error {
	requestBody := c.Locals("body").(*LoginToAccountRequestBody)

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

func (server *ApiServer) handleSetLearningPathStopBody(c *fiber.Ctx) error {
	requestBody := c.Locals("body").(*SetLearningPathStopBodyRequestBody)

	var response *SetLearningPathStopTitleResponseBody = NewSetLearningPathStopTitleResponseBody()

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "invalid table ID")
	}

	stop, err := server.store.GetLearningPathStopByID(userTableID, requestBody.LearningPathStopID)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "failed to get learning path stop")
	}

	err = server.store.SetLearningPathStopBody(stop, requestBody.Body)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "failed to update learning path stop body")
	}

	return sendSuccess(response, c)
}

func (server *ApiServer) handleSetLearningPathStopTitle(c *fiber.Ctx) error {
	requestBody := c.Locals("body").(*SetLearningPathStopTitleRequestBody)

	response := NewSetLearningPathStopTitleResponseBody()

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "invalid table ID")
	}

	stop, err := server.store.GetLearningPathStopByID(userTableID, requestBody.LearningPathStopID)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "failed to get learning path stop")
	}

	err = server.store.SetLearningPathStopTitle(stop, requestBody.Title)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "failed to update learning path stop title")
	}

	return sendSuccess(response, c)
}

func (server *ApiServer) handleSetLearningPathTitle(c *fiber.Ctx) error {
	requestBody := c.Locals("body").(*SetLearningPathTitleRequestBody)

	response := NewSetLearningPathTitleResponseBody()

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "invalid table ID")
	}

	learningPath, err := server.store.GetLearningPathByID(userTableID, requestBody.LearningPathID)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "failed to get learning path")
	}

	err = server.store.SetLearningPathTitle(learningPath, requestBody.Title)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "failed to update learning path title")
	}

	return sendSuccess(response, c)
}

func (server *ApiServer) handleCreateLearningPath(c *fiber.Ctx) error {
	requestBody := c.Locals("body").(*CreateLearningPathRequestBody)

	response := NewCreateLearningPathResponseBody()

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "invalid table ID")
	}

	_, err = server.store.CreateLearningPath(requestBody.Title, userTableID)

	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Failed to save learning path in DB")
	}

	return sendSuccess(response, c)
}

func (server *ApiServer) handleCreateLearningPathStop(c *fiber.Ctx) error {
	requestBody := c.Locals("body").(*CreateLearningPathStopRequestBody)

	var response *CreateLearningPathStopResponseBody = NewCreateLearningPathStopResponseBody()

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "invalid table ID")
	}

	// get the learning path
	learningPath, err := server.store.GetLearningPathByID(userTableID, requestBody.LearningPathID)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Failed to get learning path in DB")
	}

	// increment the next stop number
	err = server.store.IncrementStopCount(learningPath)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Failed to increment stop count")
	}

	var stop *storage.LearningPathStop = storage.NewLearningPathStop()
	stop.MarkdownBody = requestBody.MarkdownBody
	stop.StopNumber = requestBody.Stop
	stop.Title = requestBody.Title

	err = server.store.AddStopToLearningPath(learningPath, stop)
	if err != nil {
		utility.LogError(err)
		return sendInternalServerError(response, c, "Failed to save learning path stop in DB")
	}

	return sendSuccess(response, c)
}
