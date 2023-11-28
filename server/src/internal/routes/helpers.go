package routes

import (
	"github.com/adamjeanlaurent/LearningPathsApp/internal/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func sendResponse(responseBody IBaseResponseBody, c *fiber.Ctx, httpStatus int, responseCode ResponseCode, errorMessage string) error {
	responseBody.SetErrorMessage(errorMessage)
	responseBody.SetResponseCode(responseCode)
	return c.Status(httpStatus).JSON(responseBody)
}

func newBaseResponseBody() BaseResponseBody {
	return BaseResponseBody{
		ErrorMessage: "",
		ResponseCode: ResponseCode_Success,
	}
}

func parseRequestBody(out interface{}, c *fiber.Ctx) error {
	return c.BodyParser(&out)
}

func queryGetUserByEmail(db *gorm.DB, email string, user *models.User) *gorm.DB {
	return db.Where("email = ?", email).First(user)
}
