package update_user_by_id

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/cucumber/godog"
	"github.com/javiertelioz/template-clean-architecture-go/internal/application/dto/user"
	userentity "github.com/javiertelioz/template-clean-architecture-go/internal/domain/entities/user"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	featureContext := NewUserFeatureContext()
	featureContext.InitializeScenario(ctx)
}

func (ctx *UserFeatureContext) iUpdateTheUserWithID(userID string, payload *godog.DocString) error {
	var userDto user.UpdateUserDTO

	err := json.Unmarshal([]byte(payload.Content), &userDto)
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	userDto.ID = userID

	ctx.setupMocksForUserUpdate(userDto)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/%s", userID), bytes.NewBuffer([]byte(payload.Content)))
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
	body := strings.TrimSpace(ctx.responseRecorder.Body.String())
	expectedResponse = strings.TrimSpace(expectedResponse)

	body = strings.ReplaceAll(body, "\n", "")
	expectedResponse = strings.ReplaceAll(expectedResponse, "\n", "")

	if body != expectedResponse {
		return fmt.Errorf("expected response %s but got %s", expectedResponse, body)
	}

	return nil
}

// Helper functions
func (ctx *UserFeatureContext) setupMocksForUserUpdate(userDto user.UpdateUserDTO) {
	if userDto.ID == "999" {
		ctx.userRepository.On("GetByID", userDto.ID).Return(&userentity.User[string]{}, fmt.Errorf("User not found"))
	} else {
		updatedUser := userentity.NewUser[string](
			userentity.WithID[string](userDto.ID),
			userentity.WithName(userDto.Name),
			userentity.WithEmail(userDto.Email),
			userentity.WithDob[string](userDto.Dob.Time),
		)
		ctx.userRepository.On("GetByID", userDto.ID).Return(updatedUser, nil)
		ctx.userRepository.On("Update", updatedUser).Return(nil)
	}
}
