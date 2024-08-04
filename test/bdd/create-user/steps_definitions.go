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

func (ctx *UserFeatureContext) iCreateAUserWithNameAndEmail(name, email string) error {

	ctx.userRepository.On("GetByEmail", email).Return(userentity.NewUser[string](), nil)
	ctx.userRepository.On("Create", mock.Anything).Return(nil)

	userDto := user.CreateUserDTO{
		Name:  name,
		Email: email,
	}

	body, _ := json.Marshal(userDto)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	ctx.responseRecorder = httptest.NewRecorder()
	ctx.controller.CreateUser(ctx.responseRecorder, req)

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
