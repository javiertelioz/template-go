package get_all_users

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
			Paths:               []string{"get_all_users.feature"},
			Randomize:           0,
			ShowStepDefinitions: false,
			NoColors:            false,
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
