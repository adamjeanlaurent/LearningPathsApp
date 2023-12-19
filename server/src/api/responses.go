package api

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
	*BaseResponseBody
}

func NewCreateAccountResponseBody() *CreateAccountResponseBody {
	return &CreateAccountResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}

type LoginToAccountResponseBody struct {
	*BaseResponseBody
}

func NewLoginToAccountResponseBody() *LoginToAccountResponseBody {
	return &LoginToAccountResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}

type CreateLearningPathResponseBody struct {
	*BaseResponseBody
}

func NewCreateLearningPathResponseBody() *CreateLearningPathResponseBody {
	return &CreateLearningPathResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}
