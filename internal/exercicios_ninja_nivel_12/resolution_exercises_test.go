// Package exercicios_ninja_nivel_12 contains tests for the exercises of level 12 in the "aprendago" project.
// These tests are designed to validate the solutions to the practical exercises, ensuring that the output
// matches the expected results. The package utilizes various helper packages such as logger, output, and trim
// to capture and compare the output of the exercises.
package exercicios_ninja_nivel_12

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used for displaying the expected and actual output
// in a structured manner. It includes placeholders for the expected value ("want")
// and the actual value ("got"), making it easier to compare them in test results.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output and
// the strings.Contains method to check if the captured output matches the expected output.
// If there is an error capturing the output or if the output does not match the expected result,
// the test will fail and log the appropriate error message.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Error capturing output: %v", err)
		t.Fatalf("Error capturing output: %v", err)
	}

	expected := `
A idade do cachorro em anos humanos é: 21
	`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expected)) {
		t.Errorf(expectTemplate, expected, result)
	}
}
