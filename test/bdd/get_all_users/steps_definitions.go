package get_all_users

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/cucumber/godog"

	userentity "github.com/javiertelioz/template-clean-architecture-go/src/domain/entities/user"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	featureContext := NewUserFeatureContext()
	featureContext.InitializeScenario(ctx)
}

func (ctx *UserFeatureContext) iRequestTheListOfUsers() error {
	ctx.setupMocksForUserList()

	req, _ := http.NewRequest("GET", "/", nil)
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

func (ctx *UserFeatureContext) theResponseShouldBe(expectedResponse *godog.DocString) error {
	body := strings.TrimSpace(ctx.responseRecorder.Body.String())
	expected := strings.TrimSpace(expectedResponse.Content)

	body = strings.ReplaceAll(body, "\n", "")
	expected = strings.ReplaceAll(expected, "\n", "")

	if body != expected {
		return fmt.Errorf("expected response %s but got %s", expected, body)
	}

	return nil
}

func (ctx *UserFeatureContext) setupMocksForUserList() {
	users := []*userentity.User[string]{
		userentity.NewUser[string](
			userentity.WithID[string]("1"),
			userentity.WithName("Jane Doe"),
			userentity.WithEmail("jane.doe@example.com"),
		),
		userentity.NewUser[string](
			userentity.WithID[string]("2"),
			userentity.WithName("John Doe"),
			userentity.WithEmail("john.doe@example.com"),
		),
	}
	ctx.userRepository.On("GetUsers").Return(users, nil)
}
