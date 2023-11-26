package routes

import (
	"github.com/adamjeanlaurent/LearningPathsApp/internal/database/models"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/security"
	"github.com/gofiber/fiber/v2"
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
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Parsing error"
		return SendResponse(response, fiber.StatusBadRequest, c)
	}

	queryResult := handler.DB.Where("email = ?", requestBody.Email).First(nil)

	if queryResult.Error != nil {
		response.ResponseCode = ResponseCode_AccountWithEmailAlreadyExists
		return SendResponse(response, fiber.StatusOK, c)
	}

	// TODO: check for valid email

	passwordHash, err := security.HashPassword(requestBody.Password)

	if err != nil {
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Password hashing failed"
		return SendResponse(response, fiber.StatusInternalServerError, c)
	}

	user := models.User{Email: requestBody.Email, Hash: passwordHash}
	var creationResult *gorm.DB = handler.DB.Create(&user)

	if creationResult.Error != nil {
		response.ResponseCode = ResponseCode_GenericError
		response.ErrorMessage = "Failed to save user in DB"
		return SendResponse(response, fiber.StatusInternalServerError, c)
	}

	return SendResponse(response, fiber.StatusOK, c)
}
