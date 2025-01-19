// Package exercicios_ninja_nivel_6 contains test functions for exercises in level 6 of the Ninja Go course.
// These tests are designed to validate the solutions to the practical exercises provided in the course.
package exercicios_ninja_nivel_6

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used to display the expected and actual values
// in a test output. It includes placeholders for the expected value ("want") and
// the actual value ("got"), each followed by a newline for better readability.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output and
// the trim.String method to normalize the strings before comparison.
// If the captured output does not contain the expected result, the test fails
// and an error message is logged.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Inteiro: 42
Inteiro&Texto: 42 e Olá, mundo!
	`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
