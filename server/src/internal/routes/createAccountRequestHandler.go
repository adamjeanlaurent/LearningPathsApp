package routes

import (
	"github.com/adamjeanlaurent/LearningPathsApp/internal/database/models"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/logger"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/security"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (handler *RequestHandlerClient) CreateAccountRequestHandler(c *fiber.Ctx) error {
	var requestBody CreateAccountRequestBody

	response := CreateAccountResponseBody{
		BaseResponseBody: BaseResponseBody{
			ErrorMessage: "",
			ResponseCode: ResponseCode_Success,
		},
	}

	// Parse JSON from the request body
	var err error = c.BodyParser(&requestBody)
	if err != nil {
		logger.LogError(err)
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Parsing error"
		return SendResponse(&response, fiber.StatusBadRequest, c)
	}

	existingUser := &models.User{}

	queryResult := handler.DB.Where("email = ?", requestBody.Email).First(existingUser)

	if queryResult.Error != nil && queryResult.Error != gorm.ErrRecordNotFound {
		logger.LogError(queryResult.Error)
		response.ResponseCode = ResponseCode_AccountWithEmailAlreadyExists
		return SendResponse(&response, fiber.StatusOK, c)
	}

	if existingUser.Email != "" {
		response.ResponseCode = ResponseCode_AccountWithEmailAlreadyExists
		return SendResponse(&response, fiber.StatusOK, c)
	}

	// TODO: check for valid email

	passwordHash, err := security.HashPassword(requestBody.Password)

	if err != nil {
		logger.LogError(err)
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Password hashing failed"
		return SendResponse(&response, fiber.StatusInternalServerError, c)
	}

	user := models.User{
		Email: requestBody.Email,
		Hash:  passwordHash,
		BaseModel: models.BaseModel{
			StableId: uuid.New().String(),
		},
	}

	var creationResult *gorm.DB = handler.DB.Create(&user)

	if creationResult.Error != nil {
		logger.LogError(creationResult.Error)
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Failed to save user in DB"
		return SendResponse(&response, fiber.StatusInternalServerError, c)
	}

	jwtToken, err := security.CreateNewJwt(user.StableId)

	if err != nil {
		logger.LogError(err)
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Failed to generate jwt"
		return SendResponse(&response, fiber.StatusInternalServerError, c)
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	return SendResponse(&response, fiber.StatusOK, c)
}
