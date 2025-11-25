package errors

const (
	DATABASE_CONNECTION_ERROR_CODE = COMMON_ERROR_CODE + iota + 1
	INVALID_REQUEST_ERROR_CODE
)

var (
	DatabaseConnectionError = NewServerError(DATABASE_CONNECTION_ERROR_CODE, "database error occurred")
	InvalidRequestError     = NewServerError(INVALID_REQUEST_ERROR_CODE, "invalid request")
)
