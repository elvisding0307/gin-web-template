package errors

var (
	ErrParseRequest       = NewServerError(10001, "failed to parse request")
	ErrInvalidCredentials = NewServerError(10002, "invalid username or password")
	ErrUserNotFound       = NewServerError(10003, "user not found")
	ErrDatabaseConnection = NewServerError(10004, "database error occurred")
	ErrTokenGeneration    = NewServerError(10005, "token generation failed")
	ErrUserAlreadyExists  = NewServerError(10006, "user already exists")
	ErrRegistrationFailed = NewServerError(10007, "registration failed")
)
