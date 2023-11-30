package routes

import (
	"github.com/adamjeanlaurent/LearningPathsApp/internal/database/models"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/logger"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/security"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func (handler *RequestHandlerClient) createAccountRequestHandler(c *fiber.Ctx) error {
	var requestBody CreateAccountRequestBody

	response := CreateAccountResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}

	var err error = parseRequestBody(&requestBody, c)
	if err != nil || requestBody.Email == "" || requestBody.Password == "" {
		logger.LogError(err)
		return sendParsingErrorResponse(&response, c)
	}

	existingUser := models.User{}

	queryResult := queryGetUserByEmail(handler.DB, requestBody.Email, &existingUser)

	if queryResult.Error != nil && queryResult.Error != gorm.ErrRecordNotFound {
		logger.LogError(queryResult.Error)
		return sendResponse(&response, c, fiber.StatusOK, ResponseCode_AccountWithEmailAlreadyExists, "")
	}

	if existingUser.Email != "" {
		response.ResponseCode = ResponseCode_AccountWithEmailAlreadyExists
		return sendResponse(&response, c, fiber.StatusOK, ResponseCode_AccountWithEmailAlreadyExists, "")
	}

	// TODO: check for valid email

	passwordHash, err := security.HashPassword(requestBody.Password)

	if err != nil {
		logger.LogError(err)
		return sendInternalServerError(&response, c, "Password hashing failed")
	}

	user := models.User{
		Email:     requestBody.Email,
		Hash:      passwordHash,
		BaseModel: models.NewBaseModel(),
	}

	var creationResult *gorm.DB = handler.DB.Create(&user)

	if creationResult.Error != nil {
		logger.LogError(creationResult.Error)
		return sendInternalServerError(&response, c, "Failed to save user in DB")
	}

	jwtToken, err := security.CreateNewJwt(user.StableId, user.ID)

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
