package routes

import (
	"github.com/adamjeanlaurent/LearningPathsApp/internal/database/models"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/logger"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/security"
	"github.com/gofiber/fiber/v2"
)

func (handler *RequestHandlerClient) loginToAccountRequestHandler(c *fiber.Ctx) error {
	var requestBody LoginToAccountRequestBody

	response := LoginToAccountResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}

	var err error = parseRequestBody(&requestBody, c)
	if err != nil || requestBody.Email == "" || requestBody.Password == "" {
		logger.LogError(err)
		return sendResponse(&response, c, fiber.StatusBadRequest, ResponseCode_GenericError, "Parsing error")
	}

	existingUser := &models.User{}

	queryResult := queryGetUserByEmail(handler.DB, requestBody.Email, existingUser)

	if queryResult.Error != nil {
		logger.LogError(queryResult.Error)
		return sendResponse(&response, c, fiber.StatusUnauthorized, ResponseCode_NoAccountWithEmailFound, "")
	}

	err = security.ComparePasswordWithHash(requestBody.Password, existingUser.Hash)
	if err != nil {
		logger.LogError(err)
		return sendResponse(&response, c, fiber.StatusUnauthorized, ResponseCode_GenericError, "")
	}

	jwtToken, err := security.CreateNewJwt(existingUser.StableId)
	if err != nil {
		logger.LogError(err)
		return sendResponse(&response, c, fiber.StatusInternalServerError, ResponseCode_GenericError, "Failed to generate jwt")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	return sendResponse(&response, c, fiber.StatusOK, ResponseCode_Success, "")
}
