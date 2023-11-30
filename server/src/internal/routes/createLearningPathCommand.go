package routes

import (
	"github.com/adamjeanlaurent/LearningPathsApp/internal/database/models"
	"github.com/adamjeanlaurent/LearningPathsApp/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func (handler *RequestHandlerClient) CreateLearningPathRequestHandler(c *fiber.Ctx) error {
	var requestBody CreateLearningPathRequestBody

	response := CreateLearningPathResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}

	var err error = parseRequestBody(&requestBody, c)
	if err != nil || requestBody.Title == "" {
		logger.LogError(err)
		return sendParsingErrorResponse(&response, c)
	}

	userTableID, err := getUserTableIDFromContext(c)
	if err != nil {
		logger.LogError(err)
		return sendInternalServerError(&response, c, "invalid table ID")
	}

	learningPath := models.LearningPath{
		Title:     requestBody.Title,
		UserID:    userTableID,
		BaseModel: models.NewBaseModel(),
	}

	var creationResult *gorm.DB = handler.DB.Create(&learningPath)

	if creationResult.Error != nil {
		logger.LogError(creationResult.Error)
		return sendInternalServerError(&response, c, "Failed to save learning path in DB")
	}

	return sendSuccess(&response, c)
}
