// Package exercicios_ninja_nivel_8 contains tests for the exercises in level 8 of the Ninja Go course.
// These tests are designed to validate the functionality of the solutions provided for the exercises.
package exercicios_ninja_nivel_8

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used to display the expected and actual output
// in a test result. It includes placeholders for the expected value (%s) and the
// actual value (%s), and separates them with labels "want:" and "got:" for clarity.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and comparing it to the expected output.
// It uses the output.Capture method to capture the function's output and
// the trim.String method to normalize the output and expected strings before comparison.
// If the captured output does not contain the expected output, the test fails with an error message.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Exemplo 1

User 1: {"First":"M","Age":54}
User 2: {"First":"Q","Age":64}
User 3: {"First":"Moneypenny","Age":27}
User 4: {"First":"James","Age":32}

Exemplo 2

[{"First":"M","Age":54},{"First":"Q","Age":64},{"First":"Moneypenny","Age":27},{"First":"James","Age":32}]`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
