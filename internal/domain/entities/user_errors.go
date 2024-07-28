package entities

import "strings"

const (
	InvalidNameErrorCode  = 1001
	InvalidEmailErrorCode = 1002
)

type DomainError struct {
	Field       string
	Cause       string
	Message     string
	Description string
	Code        int
}

func NewDomainError(field, cause, message, description string, code int) *DomainError {
	return &DomainError{
		Field:       field,
		Cause:       cause,
		Message:     message,
		Description: description,
		Code:        code,
	}
}

func (e *DomainError) Error() string {
	return e.Message
}

type ValidationErrors struct {
	Errors []*DomainError
}

func (ve *ValidationErrors) Add(err *DomainError) {
	ve.Errors = append(ve.Errors, err)
}

func (ve *ValidationErrors) Error() string {
	var messages []string
	for _, err := range ve.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, ", ")
}

func (ve *ValidationErrors) IsEmpty() bool {
	return len(ve.Errors) == 0
}
