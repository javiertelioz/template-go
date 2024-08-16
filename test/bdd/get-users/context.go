package get_users

import (
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/controllers"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/routes"
)

type UserFeatureContext struct {
	router           *chi.Mux
	responseRecorder *httptest.ResponseRecorder
	userRepository   *repository.MockUserRepository
}

func NewUserFeatureContext() *UserFeatureContext {
	userRepository := new(repository.MockUserRepository)

	getUsersUseCase := use_cases.NewGetUsesUseCase(userRepository)

	controller := controllers.NewUserController(
		use_cases.CreateUserUseCase{},
		*getUsersUseCase,
		use_cases.GetUserByIDUseCase{},
		use_cases.UpdateUserByIDUseCase{},
		use_cases.DeleteUserByIDUseCase{},
	)

	router := routes.UserRoutes(controller)

	return &UserFeatureContext{
		router:           router.(*chi.Mux),
		responseRecorder: httptest.NewRecorder(),
		userRepository:   userRepository,
	}
}

func (ctx *UserFeatureContext) InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^I request the list of users$`, ctx.iRequestTheListOfUsers)
	s.Step(`^I should get status code (\d+)$`, ctx.iShouldGetStatusCode)
	s.Step(`^the response should be:$`, ctx.theResponseShouldBe)
}
