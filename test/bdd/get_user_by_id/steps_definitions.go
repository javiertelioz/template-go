package get_user_by_id

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	featureContext := NewUserFeatureContext()
	featureContext.InitializeScenario(ctx)
}

func (ctx *UserFeatureContext) aRepositoryWithUserNotFoundWithUserID(userID string) {
	ctx.userRepository.On("GetByID", userID).Return(&user.User[string]{}, fmt.Errorf("user not found")).Once()
}

func (ctx *UserFeatureContext) aRepositoryWithUserNotFound(userID string) {
	existingUser := user.NewUser[string](
		user.WithID(userID),
		user.WithName("John Doe"),
		user.WithEmail("john.doe@example.com"),
	)
	ctx.userRepository.On("GetByID", userID).Return(existingUser, nil).Once()
}

func (ctx *UserFeatureContext) iGetAUserByID(userID string) error {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%s", userID), nil)
	chiCtx := chi.NewRouteContext()

	request := req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	chiCtx.URLParams.Add("id", fmt.Sprintf("%v", userID))

	ctx.responseRecorder = httptest.NewRecorder()
	ctx.controller.GetUserByID(ctx.responseRecorder, request)

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
