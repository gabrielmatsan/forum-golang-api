package usecaseserror

import "fmt"

type UseCaseError struct {
	Code    string
	Message string
}

func (e *UseCaseError) Error() string {
	return e.Message
}

func NewEmailAlreadyUsedError(email string) *UseCaseError {
	return &UseCaseError{
		Code:    "EMAIL_ALREADY_USED",
		Message: fmt.Sprintf("Email \"%s\" is already in use", email),
	}
}

func NewWeakPasswordError() *UseCaseError {
	return &UseCaseError{
		Code:    "WEAK_PASSWORD",
		Message: "Password is too weak",
	}
}

func NewWrongCredentialError() *UseCaseError {
	return &UseCaseError{
		Code:    "WRONG_CREDENTIAL",
		Message: "Email or password is incorrect",
	}
}

func ResourceNotFoundError(resource string, identifier string) *UseCaseError {
	return &UseCaseError{
		Code:    "RESOURCE_NOT_FOUND",
		Message: fmt.Sprintf("%s \"%s\" not found", resource, identifier),
	}
}

func NewInvalidEmailError(email string) *UseCaseError {
	return &UseCaseError{
		Code:    "INVALID_EMAIL",
		Message: fmt.Sprintf("Email \"%s\" is invalid", email),
	}
}

func NewInternalError() *UseCaseError {
	return &UseCaseError{
		Code:    "INTERNAL_ERROR",
		Message: "An internal error occurred",
	}
}
