package routes

type ResponseCode int

const (
	ResponseCode_Success ResponseCode = iota
	ResponseCode_AccountWithEmailAlreadyExists

	ResponseCode_GenericError
)
