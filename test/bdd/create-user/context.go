package create_user

import (
	"net/http/httptest"

	"github.com/cucumber/godog"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/controllers"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type UserFeatureContext struct {
	controller            *controllers.UserController
	responseRecorder      *httptest.ResponseRecorder
	userRepository        *repository.MockUserRepository
	createUserUseCase     use_cases.CreateUserUseCase
	getUsersUseCase       use_cases.GetUsesUseCase
	getUserByIDUseCase    use_cases.GetUserByIDUseCase
	updateUserByIDUseCase use_cases.UpdateUserByIDUseCase
	deleteUserByIDUseCase use_cases.DeleteUserByIDUseCase
}

func NewUserFeatureContext() *UserFeatureContext {
	userRepository := new(repository.MockUserRepository)
	createUserUseCase := use_cases.NewCreateUserUseCase(userRepository)
	getUsersUseCase := use_cases.NewGetUsesUseCase(userRepository)
	getUserByIDUseCase := use_cases.NewGetUserByIDUseCase(userRepository)
	updateUserByIDUseCase := use_cases.NewUpdateUserByIDUseCase(userRepository)
	deleteUserByIDUseCase := use_cases.NewDeleteUserByIDUseCase(userRepository)
	controller := controllers.NewUserController(
		*createUserUseCase,
		*getUsersUseCase,
		*getUserByIDUseCase,
		*updateUserByIDUseCase,
		*deleteUserByIDUseCase,
	)

	return &UserFeatureContext{
		controller:            controller,
		responseRecorder:      httptest.NewRecorder(),
		userRepository:        userRepository,
		createUserUseCase:     *createUserUseCase,
		getUsersUseCase:       *getUsersUseCase,
		getUserByIDUseCase:    *getUserByIDUseCase,
		updateUserByIDUseCase: *updateUserByIDUseCase,
		deleteUserByIDUseCase: *deleteUserByIDUseCase,
	}
}

func (ctx *UserFeatureContext) InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^I create a user with name "([^"]*)" and email "([^"]*)"$`, ctx.iCreateAUserWithNameAndEmail)
	s.Step(`^I should get status code (\d+)$`, ctx.iShouldGetStatusCode)
	s.Step(`^the response should be "([^"]*)"$`, ctx.theResponseShouldBe)
}
