package get_user_by_id

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"

	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	featureContext := NewUserFeatureContext()
	featureContext.InitializeScenario(ctx)
}

func (ctx *UserFeatureContext) iRequestTheUserWithID(userID string) error {
	ctx.setupMocksForUserFetch(userID)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", userID), nil)
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
func (ctx *UserFeatureContext) setupMocksForUserFetch(userID string) {
	if userID == "999" {
		ctx.userRepository.On("GetByID", userID).Return(&userentity.User[string]{}, fmt.Errorf("User not found"))
	} else {
		user := userentity.NewUser[string](
			userentity.WithID[string](userID),
			userentity.WithName("Jane Doe"),
			userentity.WithEmail("jane.doe@example.com"),
		)
		ctx.userRepository.On("GetByID", userID).Return(user, nil)
	}
}
