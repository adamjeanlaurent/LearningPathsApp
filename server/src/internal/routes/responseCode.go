package routes

type ResponseCode string

const (
	ResponseCode_Success                       = "Success"
	ResponseCode_AccountWithEmailAlreadyExists = "AccountWithEmailAlreadyExists"
	ResponseCode_GenericError                  = "GenericError"
)
