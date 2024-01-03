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

func newBaseResponseBody() *BaseResponseBody {
	return &BaseResponseBody{
		ErrorMessage: "",
		ResponseCode: ResponseCode_Success,
	}
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

type CreateLearningPathStopResponseBody struct {
	*BaseResponseBody
}

func NewCreateLearningPathStopResponseBody() *CreateLearningPathStopResponseBody {
	return &CreateLearningPathStopResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}

type SetLearningPathTitleResponseBody struct {
	*BaseResponseBody
}

func NewSetLearningPathTitleResponseBody() *SetLearningPathTitleResponseBody {
	return &SetLearningPathTitleResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}

type SetLearningPathStopTitleResponseBody struct {
	*BaseResponseBody
}

func NewSetLearningPathStopTitleResponseBody() *SetLearningPathStopTitleResponseBody {
	return &SetLearningPathStopTitleResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}

type SetLearningPathStopBodyResponseBody struct {
	*BaseResponseBody
}

func NewSetLearningPathStopBodyResponseBody() *SetLearningPathStopBodyResponseBody {
	return &SetLearningPathStopBodyResponseBody{
		BaseResponseBody: newBaseResponseBody(),
	}
}
