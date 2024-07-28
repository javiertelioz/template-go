package entities

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
)

type UserTestSuite struct {
	suite.Suite
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) givenValidUserOptions() []user.UserOption[string] {
	return []user.UserOption[string]{
		user.WithID[string]("123"),
		user.WithName[string]("John Doe"),
		user.WithEmail[string]("john.doe@example.com"),
	}
}

func (suite *UserTestSuite) givenInvalidUserOptions() []user.UserOption[string] {
	return []user.UserOption[string]{
		user.WithID[string]("123"),
		user.WithName[string](""),
		user.WithEmail[string]("invalid-email"),
	}
}

// When methods
func (suite *UserTestSuite) whenCreatingUser(opts []user.UserOption[string]) *user.User[string] {
	return user.NewUser(opts...)
}

func (suite *UserTestSuite) whenValidatingUser(u *user.User[string]) *entities.ValidationErrors {
	return u.Validate()
}

// Then methods
func (suite *UserTestSuite) thenExpectNoValidationErrors(errs *entities.ValidationErrors) {
	suite.Empty(errs.Errors)
}

func (suite *UserTestSuite) thenExpectValidationErrors(errs *entities.ValidationErrors) {
	suite.NotEmpty(errs.Errors)
}

func (suite *UserTestSuite) thenExpectUserFieldsToMatch(u *user.User[string], id, name, email string) {
	suite.Equal(id, u.GetID())
	suite.Equal(name, u.GetName())
	suite.Equal(email, u.GetEmail())
}

// Test methods
func (suite *UserTestSuite) TestCreateValidUser() {
	// Given
	opts := suite.givenValidUserOptions()

	// When
	u := suite.whenCreatingUser(opts)

	// Then
	suite.thenExpectUserFieldsToMatch(u, "123", "John Doe", "john.doe@example.com")
}

func (suite *UserTestSuite) TestValidateValidUser() {
	// Given
	opts := suite.givenValidUserOptions()
	u := suite.whenCreatingUser(opts)

	// When
	errs := suite.whenValidatingUser(u)

	// Then
	suite.thenExpectNoValidationErrors(errs)
}

func (suite *UserTestSuite) TestValidateInvalidUser() {
	// Given
	opts := suite.givenInvalidUserOptions()
	u := suite.whenCreatingUser(opts)

	// When
	errs := suite.whenValidatingUser(u)

	// Then
	suite.thenExpectValidationErrors(errs)
}
