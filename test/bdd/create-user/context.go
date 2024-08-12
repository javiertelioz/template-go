package create_user

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
	// Inicializamos el mock del repositorio de usuarios
	userRepository := new(repository.MockUserRepository)

	// Creamos los casos de uso utilizando el repositorio mockeado
	createUserUseCase := use_cases.NewCreateUserUseCase(userRepository)
	getUsersUseCase := use_cases.NewGetUsesUseCase(userRepository)
	getUserByIDUseCase := use_cases.NewGetUserByIDUseCase(userRepository)
	updateUserByIDUseCase := use_cases.NewUpdateUserByIDUseCase(userRepository)
	deleteUserByIDUseCase := use_cases.NewDeleteUserByIDUseCase(userRepository)

	// Creamos el controlador usando los casos de uso creados
	controller := controllers.NewUserController(
		*createUserUseCase,
		*getUsersUseCase,
		*getUserByIDUseCase,
		*updateUserByIDUseCase,
		*deleteUserByIDUseCase,
	)

	// Configuramos el router con los handlers
	router := routes.UserRoutes(controller)

	return &UserFeatureContext{
		router:           router.(*chi.Mux),
		responseRecorder: httptest.NewRecorder(),
		userRepository:   userRepository,
	}
}

func (ctx *UserFeatureContext) InitializeScenario(s *godog.ScenarioContext) {
	// Registramos los pasos definidos en el archivo steps_definitions.go
	s.Step(`^I create a user with payload:$`, ctx.iCreateAUserWithPayload)
	s.Step(`^I should get status code (\d+)$`, ctx.iShouldGetStatusCode)
	s.Step(`^the response should be "([^"]*)"$`, ctx.theResponseShouldBe)
}
