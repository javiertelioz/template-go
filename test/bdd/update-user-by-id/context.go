package update_user_by_id

import (
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/go-chi/chi/v5"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/use_cases"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/controllers"
	"github.com/javiertelioz/template-clean-architecture-go/internal/presentation/routes"
	"github.com/javiertelioz/template-clean-architecture-go/test/mocks/repository"
)

type UserFeatureContext struct {
	router           *chi.Mux
	responseRecorder *httptest.ResponseRecorder
	userRepository   *repository.MockUserRepository
}

func NewUserFeatureContext() *UserFeatureContext {
	userRepository := new(repository.MockUserRepository)

	updateUserByIDUseCase := use_cases.NewUpdateUserByIDUseCase(userRepository)

	controller := controllers.NewUserController(
		use_cases.CreateUserUseCase{},
		use_cases.GetUsesUseCase{},
		use_cases.GetUserByIDUseCase{},
		*updateUserByIDUseCase,
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
	s.Step(`^I update the user with ID "([^"]*)" with payload:$`, ctx.iUpdateTheUserWithID)
	s.Step(`^I should get status code (\d+)$`, ctx.iShouldGetStatusCode)
	s.Step(`^the response should be "([^"]*)"$`, ctx.theResponseShouldBe)
}
