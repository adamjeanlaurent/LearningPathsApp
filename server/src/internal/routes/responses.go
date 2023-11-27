package routes

import "github.com/gofiber/fiber/v2"

func SendResponse(responseBody IBaseResponseBody, c *fiber.Ctx, httpStatus int, responseCode ResponseCode, errorMessage string) error {
	responseBody.SetErrorMessage(errorMessage)
	responseBody.SetResponseCode(responseCode)
	return c.Status(httpStatus).JSON(responseBody)
}

func NewBaseResponseBody() BaseResponseBody {
	return BaseResponseBody{
		ErrorMessage: "",
		ResponseCode: ResponseCode_Success,
	}
}

type IBaseResponseBody interface {
	SetResponseCode(code ResponseCode)
	SetErrorMessage(message string)
}

type BaseResponseBody struct {
	ErrorMessage string `json:"errorMessage"`
	ResponseCode
}

func (r *BaseResponseBody) SetResponseCode(code ResponseCode) {
	r.ResponseCode = code
}

func (r *BaseResponseBody) SetErrorMessage(message string) {
	r.ErrorMessage = message
}

type CreateAccountResponseBody struct {
	BaseResponseBody
	Email string
}
