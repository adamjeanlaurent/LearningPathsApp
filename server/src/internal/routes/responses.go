package routes

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
}

type LoginToAccountResponseBody struct {
	BaseResponseBody
}

type CreateLearningPathResponseBody struct {
	BaseResponseBody
}
