// Package exercicios_ninja_nivel_4 contains solutions and tests for exercises
// aimed at practicing and improving Go programming skills. The exercises cover
// various topics and concepts, providing hands-on experience and reinforcing
// understanding of the language.
package exercicios_ninja_nivel_4

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// testTemplate is a constant string used for formatting test output.
// It provides a template for displaying the expected and actual values
// in test results, making it easier to compare them.
const (
	expectTemplate = "\nwant:\n%s\n, \ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the functionality of iterating over an array of integers,
// concatenating each integer to a string with a comma and space separator, and comparing the result
// to an expected string. If the resulting string does not match the expected string, the test fails.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	capturer := output.Capture()
	result := capturer.New(ResolucaoNaPraticaExercicio1)

	expect := `
Resolução:
1, 2, 3, 4, 5
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected output. It uses a
// deferred function to recover from any panic that might occur during the execution
// of ResolucaoNaPraticaExercicio2 and captures the panic message if it happens.
// The test checks if the captured output contains the expected string, ignoring
// leading and trailing whitespace, and reports an error if it does not match.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	capturer := output.Capture()
	result := capturer.New(ResolucaoNaPraticaExercicio2)

	expect := `
Resolução:

Slice: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
Tipo: []string
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected result. It uses the
// output.Capture utility to capture the function's output and the trim.New utility
// to normalize the strings before comparison. If the captured output does not
// contain the expected result, the test fails with an error message.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	capturer := output.Capture()
	result := capturer.New(ResolucaoNaPraticaExercicio3)

	expect := `
Resolução:

Slice: 1, 2, 3, 4
Tipo: []string
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4
// by capturing its output and comparing it to the expected result. The test
// verifies that the function correctly appends elements to slices and produces
// the expected output. If the actual output does not match the expected output,
// the test will fail and report the discrepancy.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	capturer := output.Capture()
	result := capturer.New(ResolucaoNaPraticaExercicio4)

	expect := `
Resolução:
append52: [42 43 44 45 46 47 48 49 50 51 52]
append53to55: [42 43 44 45 46 47 48 49 50 51 52 53 54 55]
appendSliceY: [42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
