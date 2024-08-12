package create_user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/mock"

	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	featureContext := NewUserFeatureContext()
	featureContext.InitializeScenario(ctx)
}

func (ctx *UserFeatureContext) iCreateAUserWithPayload(payload *godog.DocString) error {
	var userDto user.CreateUserDTO
	err := json.Unmarshal([]byte(payload.Content), &userDto)
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	// Configurar mocks en función del contenido del DTO
	ctx.setupMocksForUserCreation(userDto)

	// Ejecutar la solicitud para crear el usuario
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(payload.Content)))
	req.Header.Set("Content-Type", "application/json")

	ctx.responseRecorder = httptest.NewRecorder()
	ctx.router.ServeHTTP(ctx.responseRecorder, req)

	return nil
}

func (ctx *UserFeatureContext) iShouldGetStatusCode(expectedCode int) error {
	if status := ctx.responseRecorder.Code; status != expectedCode {
		return fmt.Errorf("expected status code %d but got %d", expectedCode, status)
	}

	return nil
}

func (ctx *UserFeatureContext) theResponseShouldBe(expectedResponse string) error {
	body := ctx.responseRecorder.Body.String()
	if body != expectedResponse {
		return fmt.Errorf("expected response %s but got %s", expectedResponse, body)
	}

	return nil
}

// Helper functions
func (ctx *UserFeatureContext) setupMocksForUserCreation(userDto user.CreateUserDTO) {
	if userDto.Name == "" || userDto.Email == "john.doeexample.com" {
		ctx.userRepository.On("GetByEmail", userDto.Email).Return(nil, nil)
	} else {
		ctx.userRepository.On("GetByEmail", userDto.Email).Return(&userentity.User[string]{}, nil)
		ctx.userRepository.On("Create", mock.Anything).Return(nil)
	}
}
