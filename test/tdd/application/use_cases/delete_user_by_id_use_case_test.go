package use_cases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type DeleteUserByIDUseCaseTestSuite struct {
	suite.Suite
	useCase        *use_cases.DeleteUserByIDUseCase
	userRepository *repository.MockUserRepository
	err            error
}

func TestDeleteUserByIDUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserByIDUseCaseTestSuite))
}

func (suite *DeleteUserByIDUseCaseTestSuite) SetupTest() {
	suite.userRepository = new(repository.MockUserRepository)
	suite.useCase = use_cases.NewDeleteUserByIDUseCase(suite.userRepository)
}

func (suite *DeleteUserByIDUseCaseTestSuite) givenUserID() string {
	return "e17d5135-989e-4977-99ef-495c0ab7cd00"
}

func (suite *DeleteUserByIDUseCaseTestSuite) givenNonExistentUserID() string {
	return "non-existent-id"
}

func (suite *DeleteUserByIDUseCaseTestSuite) whenDeleteUserByIDUseCaseIsCalled(id string) {
	suite.err = suite.useCase.Execute(id)
}

// Test methods
func (suite *DeleteUserByIDUseCaseTestSuite) TestDeleteUserByIDUseCaseWithValidIDWhenCallExecuteThenReturnSuccessResult() {
	// Given
	id := suite.givenUserID()
	existingUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
	suite.userRepository.On("GetByID", id).Return(existingUser, nil)
	suite.userRepository.On("Delete", id).Return(nil)

	// When
	suite.whenDeleteUserByIDUseCaseIsCalled(id)

	// Then
	suite.NoError(suite.err)
	suite.userRepository.AssertCalled(suite.T(), "GetByID", id)
	suite.userRepository.AssertCalled(suite.T(), "Delete", id)
}

func (suite *DeleteUserByIDUseCaseTestSuite) TestDeleteUserByIDUseCaseWithNonExistentIDWhenCallExecuteThenReturnError() {
	// Given
	id := suite.givenNonExistentUserID()
	suite.userRepository.On("GetByID", id).Return(userentity.NewUser[string](), errors.New("user not found"))

	// When
	suite.whenDeleteUserByIDUseCaseIsCalled(id)

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "user not found")
	suite.userRepository.AssertCalled(suite.T(), "GetByID", id)
	suite.userRepository.AssertNotCalled(suite.T(), "Delete", id)
}

func (suite *DeleteUserByIDUseCaseTestSuite) TestDeleteUserByIDUseCaseWithValidIDWhenCallExecuteThenReturnRepositoryError() {
	// Given
	id := suite.givenUserID()
	existingUser := userentity.NewUser(
		userentity.WithID(id),
		userentity.WithName("John Doe"),
		userentity.WithEmail("john.doe@example.com"),
	)
	suite.userRepository.On("GetByID", id).Return(existingUser, nil)
	suite.userRepository.On("Delete", id).Return(errors.New("repository error"))

	// When
	suite.whenDeleteUserByIDUseCaseIsCalled(id)

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "repository error")
	suite.userRepository.AssertCalled(suite.T(), "GetByID", id)
	suite.userRepository.AssertCalled(suite.T(), "Delete", id)
}
