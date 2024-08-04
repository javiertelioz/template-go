package user

import (
	"regexp"
	"strings"

	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities"
)

const emailRegex = `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`

var compiledEmailRegex = regexp.MustCompile(emailRegex)

func ValidateName(name string) *entities.DomainError {
	if strings.TrimSpace(name) == "" {
		return entities.NewDomainError(
			"Name",
			"empty string",
			"Name cannot be empty",
			"An empty string is not allowed for Name",
			entities.InvalidNameErrorCode,
		)
	}
	return nil
}

func ValidateEmail(email string) *entities.DomainError {
	if !compiledEmailRegex.MatchString(email) {
		return entities.NewDomainError(
			"Email",
			"invalid format",
			"Email format is invalid",
			"The email format does not match the required pattern",
			entities.InvalidEmailErrorCode,
		)
	}

	return nil
}
