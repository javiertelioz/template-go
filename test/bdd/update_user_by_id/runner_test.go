package update_user_by_id

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:              "pretty",
			Output:              colors.Colored(os.Stdout),
			Paths:               []string{"update_user_by_id.feature"},
			Randomize:           -1,
			ShowStepDefinitions: false,
			NoColors:            false,
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
