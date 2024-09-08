package create_user

import (
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"

	"github.com/javiertelioz/template-clean-architecture-go/src/application/use_cases"
	"github.com/javiertelioz/template-clean-architecture-go/src/interfaces/controllers"
	"github.com/javiertelioz/template-clean-architecture-go/src/interfaces/routes"
)

type UserFeatureContext struct {
	router           *chi.Mux
	responseRecorder *httptest.ResponseRecorder
	userRepository   *repository.MockUserRepository
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

	router := routes.UserRoutes(controller)

	return &UserFeatureContext{
		router:           router.(*chi.Mux),
		responseRecorder: httptest.NewRecorder(),
		userRepository:   userRepository,
	}
}

func (ctx *UserFeatureContext) InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^I create a user with payload:$`, ctx.iCreateAUserWithPayload)
	s.Step(`^I should get status code (\d+)$`, ctx.iShouldGetStatusCode)
	s.Step(`^the response should be "([^"]*)"$`, ctx.theResponseShouldBe)
}
