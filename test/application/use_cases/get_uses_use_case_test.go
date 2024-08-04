package use_cases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type GetUsesUseCaseTestSuite struct {
	suite.Suite
	useCase        *use_cases.GetUsesUseCase
	userRepository *repository.MockUserRepository
	result         []user.GetUserDTO
	err            error
}

func TestGetUsesUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetUsesUseCaseTestSuite))
}

func (suite *GetUsesUseCaseTestSuite) SetupTest() {
	suite.userRepository = new(repository.MockUserRepository)
	suite.useCase = use_cases.NewGetUsesUseCase(suite.userRepository)
}

func (suite *GetUsesUseCaseTestSuite) whenGetUsesUseCaseIsCalled() {
	suite.result, suite.err = suite.useCase.Execute()
}

// Test methods
func (suite *GetUsesUseCaseTestSuite) TestGetUsesUseCaseWhenCallExecuteThenReturnSuccessResult() {
	// Given
	domainUsers := []*userentity.User[string]{
		userentity.NewUser(
			userentity.WithID("1"),
			userentity.WithName("John Doe"),
			userentity.WithEmail("john.doe@example.com"),
		),
		userentity.NewUser(
			userentity.WithID("2"),
			userentity.WithName("Jane Doe"),
			userentity.WithEmail("jane.doe@example.com"),
		),
	}
	suite.userRepository.On("GetUsers").Return(domainUsers, nil)

	// When
	suite.whenGetUsesUseCaseIsCalled()

	// Then
	suite.NoError(suite.err)
	suite.Len(suite.result, 2)
	suite.Equal("1", suite.result[0].ID)
	suite.Equal("John Doe", suite.result[0].Name)
	suite.Equal("john.doe@example.com", suite.result[0].Email)
	suite.Equal("2", suite.result[1].ID)
	suite.Equal("Jane Doe", suite.result[1].Name)
	suite.Equal("jane.doe@example.com", suite.result[1].Email)
	suite.userRepository.AssertCalled(suite.T(), "GetUsers")
}

func (suite *GetUsesUseCaseTestSuite) TestGetUsesUseCaseWhenCallExecuteThenReturnError() {
	// Given
	var users []*userentity.User[string]
	suite.userRepository.On("GetUsers").Return(users, errors.New("repository error"))

	// When
	suite.whenGetUsesUseCaseIsCalled()

	// Then
	suite.Error(suite.err)
	suite.EqualError(suite.err, "repository error")
	suite.Nil(suite.result)
	suite.userRepository.AssertCalled(suite.T(), "GetUsers")
}
