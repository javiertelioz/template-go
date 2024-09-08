package use_cases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/template-clean-architecture-go/src/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/src/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type GetUserByIDUseCaseTestSuite struct {
	suite.Suite
	useCase        *use_cases.GetUserByIDUseCase
	userRepository *repository.MockUserRepository
	result         user.GetUserDTO
	err            error
}

func TestGetUserByIDUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserByIDUseCaseTestSuite))
}

func (suite *GetUserByIDUseCaseTestSuite) SetupTest() {
	suite.userRepository = new(repository.MockUserRepository)
	suite.useCase = use_cases.NewGetUserByIDUseCase(suite.userRepository)
}

func (suite *GetUserByIDUseCaseTestSuite) givenUserID() string {
	return "e17d5135-989e-4977-99ef-495c0ab7cd00"
}

func (suite *GetUserByIDUseCaseTestSuite) givenNonExistentUserID() string {
	return "non-existent-id"
}

func (suite *GetUserByIDUseCaseTestSuite) whenGetUserByIDUseCaseIsCalled(id string) {
	suite.result, suite.err = suite.useCase.Execute(id)
}

// Test methods
func (suite *GetUserByIDUseCaseTestSuite) TestGetUserByIDUseCaseWithValidIDWhenCallExecuteThenReturnSuccessResult() {
	// Given
	id := suite.givenUserID()
	domainUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
	suite.userRepository.On("GetByID", id).Return(domainUser, nil)

	// When
	suite.whenGetUserByIDUseCaseIsCalled(id)

	// Then
	suite.NoError(suite.err)
	suite.Equal(id, suite.result.ID)
	suite.Equal("John Doe", suite.result.Name)
	suite.Equal("john.doe@example.com", suite.result.Email)
	suite.userRepository.AssertCalled(suite.T(), "GetByID", id)
}

func (suite *GetUserByIDUseCaseTestSuite) TestGetUserByIDUseCaseWithNonExistentIDWhenCallExecuteThenReturnError() {
	// Given
	id := suite.givenNonExistentUserID()
	suite.userRepository.On("GetByID", id).
		Return(userentity.NewUser[string](), errors.New("user not found"))

	// When
	suite.whenGetUserByIDUseCaseIsCalled(id)

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "user not found")
	suite.userRepository.AssertCalled(suite.T(), "GetByID", id)
}
