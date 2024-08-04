package use_cases

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type UpdateUserByIDUseCaseTestSuite struct {
	suite.Suite
	useCase        *use_cases.UpdateUserByIDUseCase
	userRepository *repository.MockUserRepository
	err            error
}

func TestUpdateUserByIDUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateUserByIDUseCaseTestSuite))
}

func (suite *UpdateUserByIDUseCaseTestSuite) SetupTest() {
	suite.userRepository = new(repository.MockUserRepository)
	suite.useCase = use_cases.NewUpdateUserByIDUseCase(suite.userRepository)
}

func (suite *UpdateUserByIDUseCaseTestSuite) givenValidUpdateUserDTO() user.UpdateUserDTO {
	return user.UpdateUserDTO{
		ID:    "1",
		Name:  "John Doe Updated",
		Email: "john.doe.updated@example.com",
		Dob: user.DateOfBirth{
			Time: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
}

func (suite *UpdateUserByIDUseCaseTestSuite) givenInvalidUpdateUserDTO() user.UpdateUserDTO {
	return user.UpdateUserDTO{
		ID:    "1",
		Name:  "",
		Email: "invalid-email",
		Dob: user.DateOfBirth{
			Time: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
}

func (suite *UpdateUserByIDUseCaseTestSuite) whenUpdateUserByIDUseCaseIsCalled(dto user.UpdateUserDTO) {
	suite.err = suite.useCase.Execute(dto)
}

// Test methods
func (suite *UpdateUserByIDUseCaseTestSuite) TestUpdateUserByIDUseCaseWithValidUserWhenCallExecuteThenReturnSuccessResult() {
	// Given
	dto := suite.givenValidUpdateUserDTO()
	existingUser := userentity.NewUser(
		userentity.WithID(dto.ID),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
		userentity.WithDob[string](time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
	)
	suite.userRepository.On("GetByID", dto.ID).Return(existingUser, nil)
	suite.userRepository.On("Update", mock.Anything).Return(nil)

	// When
	suite.whenUpdateUserByIDUseCaseIsCalled(dto)

	// Then
	suite.NoError(suite.err)
	suite.userRepository.AssertCalled(suite.T(), "Update", mock.Anything)
}

func (suite *UpdateUserByIDUseCaseTestSuite) TestUpdateUserByIDUseCaseWithNonExistentIDWhenCallExecuteThenReturnError() {
	// Given
	dto := suite.givenValidUpdateUserDTO()
	suite.userRepository.On("GetByID", dto.ID).Return(userentity.NewUser[string](), errors.New("user not found"))

	// When
	suite.whenUpdateUserByIDUseCaseIsCalled(dto)

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "user not found")
	suite.userRepository.AssertNotCalled(suite.T(), "Update", mock.Anything)
}

func (suite *UpdateUserByIDUseCaseTestSuite) TestUpdateUserByIDUseCaseWithInvalidUserWhenCallExecuteThenReturnValidationError() {
	// Given
	dto := suite.givenInvalidUpdateUserDTO()
	existingUser := userentity.NewUser(
		userentity.WithID(dto.ID),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
		userentity.WithDob[string](time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
	)
	suite.userRepository.On("GetByID", dto.ID).Return(existingUser, nil)

	// When
	suite.whenUpdateUserByIDUseCaseIsCalled(dto)

	// Then
	suite.Error(suite.err)
	suite.Contains(suite.err.Error(), "Name cannot be empty")
	suite.Contains(suite.err.Error(), "Email format is invalid")
	suite.userRepository.AssertNotCalled(suite.T(), "Update", mock.Anything)
}

func (suite *UpdateUserByIDUseCaseTestSuite) TestUpdateUserByIDUseCaseWithValidUserWhenCallExecuteThenReturnRepositoryError() {
	// Given
	dto := suite.givenValidUpdateUserDTO()
	existingUser := userentity.NewUser(
		userentity.WithID(dto.ID),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
		userentity.WithDob[string](time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
	)
	suite.userRepository.On("GetByID", dto.ID).Return(existingUser, nil)
	suite.userRepository.On("Update", mock.Anything).Return(errors.New("repository error"))

	// When
	suite.whenUpdateUserByIDUseCaseIsCalled(dto)

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "repository error")
	suite.userRepository.AssertCalled(suite.T(), "Update", mock.Anything)
}
