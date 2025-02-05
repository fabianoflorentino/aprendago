// Package exercicios_ninja_nivel_10 contains tests for the exercises in level 10 of the Ninja Go course.
// These tests are designed to validate the solutions to the exercises by comparing the actual output
// of the functions to the expected output.
package exercicios_ninja_nivel_10

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used for displaying the expected and actual output
// in a structured way. It helps in comparing the expected result with the actual result
// by showing them side by side in the output.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and comparing it to the expected value.
// It fails the test if the output does not contain the expected value or if there is an error capturing the output.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Value: 42
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected result. If the output
// does not match the expected result, the test fails with an error message.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Value: 42
--------------------
Type: chan<- int
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected result. If the output
// does not match the expected result, the test fails with an error message.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Value: 0
Value: 1
Value: 2
Value: 3
Value: 4
Value: 5
Value: 6
Value: 7
Value: 8
Value: 9
Value: 10
Value: 11
Value: 12
Value: 13
Value: 14
Value: 15
Value: 16
Value: 17
Value: 18
Value: 19
Value: 20
Value: 21
Value: 22
Value: 23
Value: 24
Value: 25
Value: 26
Value: 27
Value: 28
Value: 29
Value: 30
Value: 31
Value: 32
Value: 33
Value: 34
Value: 35
Value: 36
Value: 37
Value: 38
Value: 39
Value: 40
Value: 41
Value: 42
Value: 43
Value: 44
Value: 45
Value: 46
Value: 47
Value: 48
Value: 49
Value: 50
Value: 51
Value: 52
Value: 53
Value: 54
Value: 55
Value: 56
Value: 57
Value: 58
Value: 59
Value: 60
Value: 61
Value: 62
Value: 63
Value: 64
Value: 65
Value: 66
Value: 67
Value: 68
Value: 69
Value: 70
Value: 71
Value: 72
Value: 73
Value: 74
Value: 75
Value: 76
Value: 77
Value: 78
Value: 79
Value: 80
Value: 81
Value: 82
Value: 83
Value: 84
Value: 85
Value: 86
Value: 87
Value: 88
Value: 89
Value: 90
Value: 91
Value: 92
Value: 93
Value: 94
Value: 95
Value: 96
Value: 97
Value: 98
Value: 99
About to exit...
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4
// by capturing its output and comparing it to the expected output.
// It logs and fails the test if there is an error capturing the output
// or if the captured output does not contain the expected string.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio4)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
About to exit...
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
