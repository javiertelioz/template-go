package entities

import (
	"testing"
	"time"

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
		user.WithDob[string](time.Now().UTC().AddDate(-30, 0, 0)),
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
	suite.False(errs.IsEmpty())
	suite.NotEmpty(errs.Errors)
	suite.Equal(len(errs.Errors), 2)
	suite.Contains(errs.Error(), "Name cannot be empty, Email format is invalid")
}

func (suite *UserTestSuite) thenExpectUserFieldsToMatch(u *user.User[string], id, name, email, dob string) {
	suite.Equal(id, u.GetID())
	suite.Equal(name, u.GetName())
	suite.Equal(email, u.GetEmail())
	suite.Contains(u.GetDob().String(), dob)
}

// Test methods
func (suite *UserTestSuite) TestCreateValidUser() {
	// Given
	opts := suite.givenValidUserOptions()

	// When
	u := suite.whenCreatingUser(opts)

	// Then
	suite.thenExpectUserFieldsToMatch(u, "123", "John Doe", "john.doe@example.com", "1994-08-03")
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
