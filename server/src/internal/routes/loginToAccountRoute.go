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
		return sendParsingErrorResponse(&response, c)
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

	jwtToken, err := security.CreateNewJwt(existingUser.StableId, existingUser.ID)
	if err != nil {
		logger.LogError(err)
		return sendInternalServerError(&response, c, "Failed to generate Jwt")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	return sendSuccess(&response, c)
}
