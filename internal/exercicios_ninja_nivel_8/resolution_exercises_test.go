// Package exercicios_ninja_nivel_8 contains a series of test functions designed to validate
// the output of various exercises. These exercises are part of a learning module aimed at
// practicing Go programming skills. Each test function captures the output of a specific
// exercise function and compares it to the expected output. If the captured output does not
// match the expected output, the test fails and logs an error. The package utilizes helper
// functions from other packages such as logger, output, and trim to facilitate output
// capturing, logging, and string normalization.
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
// by capturing its output and comparing it with the expected output.
// It uses the output.Capture method to capture the function's output and
// the trim.String method to normalize the strings before comparison.
// If the captured output does not contain the expected output, the test fails.
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

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected result.
// It logs and fails the test if there is an error capturing the output
// or if the output does not contain the expected string.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
User 1: First: M Age: 54
User 2: First: Q Age: 64
User 3: First: Moneypenny Age: 27
User 4: First: James Age: 32
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected result.
// It logs and fails the test if there is an error capturing the output
// or if the output does not contain the expected string.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation"]}
{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to","take care of that for you, James?"]}
{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't","have to bring it back"]}
{"First":"Q","Last":"See","Age":64,"Sayings":["Have some of the cake, James","It's all gone, because you were late"]}
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4 by capturing its output
// and comparing it against the expected output. It verifies that the function correctly sorts slices of integers
// and strings. If the captured output does not match the expected output, the test fails and logs an error.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio4)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Slice de inteiros desordenado:  [5 8 2 43 17 987 14 12 21 1 4 2 3 93 13]
Slice de inteiros ordenado:  [1 2 2 3 4 5 8 12 13 14 17 21 43 93 987]
Slice de strings desordenado:  [random rainbow delights in torpedo summers under gallantry fragmented moons across magenta]
Slice de strings ordenado:  [across delights fragmented gallantry in magenta moons rainbow random summers torpedo under]
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio5 tests the function ResolucaoNaPraticaExercicio5
// by capturing its output and comparing it to the expected result.
// It verifies that the function correctly sorts a slice of users by their first name and age.
// If the captured output does not match the expected output, the test fails.
func TestResolucaoNaPraticaExercicio5(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio5)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
By First
Slice de users desordenado:  [{M 54} {Q 64} {Moneypenny 27} {James 32}]
Slice de users ordenado por First:  [{James 32} {M 54} {Moneypenny 27} {Q 64}]

By Age
Slice de users desordenado:  [{M 54} {Q 64} {Moneypenny 27} {James 32}]
Slice de users ordenado por Age:  [{Moneypenny 27} {James 32} {M 54} {Q 64}]
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
