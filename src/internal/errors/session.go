package errors

const (
	USER_NOT_FOUND_ERROR_CODE = SESSION_ERROR_CODE + iota + 1
	TOKEN_GENERATION_ERROR_CODE
	USER_ALREADY_EXISTS_ERROR_CODE
	INVALID_CREDENTIALS_ERROR_CODE
	REGISTRATION_FAILED_ERROR_CODE
)

var (
	UserNotFoundError       = NewServerError(USER_NOT_FOUND_ERROR_CODE, "user not found")
	TokenGenerationError    = NewServerError(TOKEN_GENERATION_ERROR_CODE, "token generation failed")
	UserAlreadyExistsError  = NewServerError(USER_ALREADY_EXISTS_ERROR_CODE, "user already exists")
	InvalidCredentialsError = NewServerError(INVALID_CREDENTIALS_ERROR_CODE, "invalid username or password")
	RegistrationFailedError = NewServerError(REGISTRATION_FAILED_ERROR_CODE, "registration failed")
)
