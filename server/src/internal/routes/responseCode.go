package routes

type ResponseCode string

const (
	ResponseCode_Success = "Success"

	ResponseCode_AccountWithEmailAlreadyExists = "AccountWithEmailAlreadyExists"
	ResponseCode_NoAccountWithEmailFound       = "NoAccountWithEmailFound"

	ResponseCode_GenericError = "GenericError"
)
