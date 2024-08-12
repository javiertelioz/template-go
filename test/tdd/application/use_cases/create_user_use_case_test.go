package use_cases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	useCase        *use_cases.CreateUserUseCase
	userRepository *repository.MockUserRepository
	user           *userentity.User[string]
	err            error
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}

func (suite *CreateUserUseCaseTestSuite) SetupTest() {
	suite.userRepository = new(repository.MockUserRepository)
	suite.useCase = use_cases.NewCreateUserUseCase(suite.userRepository)

	suite.user = userentity.NewUser(
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
}

func (suite *CreateUserUseCaseTestSuite) givenValidUserDTO() user.CreateUserDTO {
	return user.CreateUserDTO{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
}

func (suite *CreateUserUseCaseTestSuite) givenInvalidUserDTO() user.CreateUserDTO {
	return user.CreateUserDTO{
		Name:  "",
		Email: "invalid-email",
	}
}

func (suite *CreateUserUseCaseTestSuite) whenCreateUserUseCaseIsCall(dto user.CreateUserDTO) {
	suite.err = suite.useCase.Execute(dto)
}

// Test methods
func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithValidUserWhenCallExecuteThenReturnSuccessResult() {
	// Given
	dto := suite.givenValidUserDTO()

	suite.userRepository.On("GetByEmail", dto.Email).Return(userentity.NewUser[string](), nil)
	suite.userRepository.On("Create", suite.user).Return(nil)

	// When
	suite.whenCreateUserUseCaseIsCall(dto)

	// Then
	suite.NoError(suite.err)
	suite.userRepository.AssertCalled(suite.T(), "Create", mock.Anything)
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithExistingEmailWhenCallExecuteThenReturnErrorOnGetUserByEmail() {
	// Given
	dto := suite.givenValidUserDTO()
	existingUser := userentity.NewUser(
		userentity.WithID("e17d5135-989e-4977-99ef-495c0ab7cd00"),
		userentity.WithName("Jane Doe"),
		userentity.WithEmail(dto.Email),
	)
	suite.userRepository.On("GetByEmail", dto.Email).Return(existingUser, nil)

	// When
	suite.whenCreateUserUseCaseIsCall(dto)

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "email already exists")
	suite.userRepository.AssertNotCalled(suite.T(), "Create", mock.Anything)
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithInvalidUserWhenCallExecuteThenReturnErrorOnValidateUser() {
	// Given
	dto := suite.givenInvalidUserDTO()
	suite.userRepository.On("GetByEmail", dto.Email).Return(userentity.NewUser[string](), nil)

	// When
	suite.whenCreateUserUseCaseIsCall(dto)

	// Then
	suite.Error(suite.err)
	suite.Contains(suite.err.Error(), "Name cannot be empty")
	suite.Contains(suite.err.Error(), "Email format is invalid")
	suite.userRepository.AssertNotCalled(suite.T(), "Create", mock.Anything)
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithValidUserWhenCallExecuteThenReturnErrorOnGetUserByEmail() {
	// Given
	dto := suite.givenValidUserDTO()
	suite.userRepository.On("GetByEmail", dto.Email).Return(userentity.NewUser[string](), errors.New("provider error"))

	// When
	suite.whenCreateUserUseCaseIsCall(dto)

	// Then
	suite.Error(suite.err)
	suite.Contains(suite.err.Error(), "provider error")
	suite.userRepository.AssertCalled(suite.T(), "GetByEmail", dto.Email)
}
