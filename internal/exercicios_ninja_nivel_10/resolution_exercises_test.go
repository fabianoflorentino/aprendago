// Package exercicios_ninja_nivel_10 contains tests for the exercises in level 10 of the Ninja Go course.
// These tests are designed to validate the solutions to the exercises by comparing the actual output
// of the functions to the expected output.
package exercicios_ninja_nivel_10

import (
	"strconv"
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

	expectValueCount := 100
	expectExitText := "About to exit..."

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expectExitText)) {
		t.Errorf(expectTemplate, expectExitText, result)
	}

	if strings.Count(trim.String(result), "Value:") != expectValueCount {
		t.Errorf("expected %d occurrences of 'Value:', got %d", expectValueCount, strings.Count(trim.String(result), "Value:"))
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

// TestResolucaoNaPraticaExercicio5 tests the function ResolucaoNaPraticaExercicio5
// by capturing its output and comparing it to the expected result.
// It verifies that the function produces the correct output and handles
// the channel state as expected.
func TestResolucaoNaPraticaExercicio5(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio5)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Value: 42
Channel open: true
Value: 0
Channel open: false
`

	trim := trim.New()

	expectChannelOpen := strings.Contains(trim.String(result), "Channel open: true")
	expectChannelClosed := strings.Contains(trim.String(result), "Channel open: false")

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}

	if !expectChannelOpen || !expectChannelClosed {
		t.Errorf("expected channel open and closed states, got open: %v, closed: %v", expectChannelOpen, expectChannelClosed)
	}
}

// TestResolucaoNaPraticaExercicio6 tests the function ResolucaoNaPraticaExercicio6
// by capturing its output and verifying the expected number of occurrences of
// the string "Channel:" and the presence of the string "Channel: 100".
// It uses the output.Capture method to capture the function's output and
// the strings.Count and strings.Contains functions to perform the checks.
// If the captured output does not meet the expectations, the test will fail
// with an appropriate error message.
func TestResolucaoNaPraticaExercicio6(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio6)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expectChannelCount := 101

	trim := trim.New()

	if strings.Count(trim.String(result), "Channel:") != expectChannelCount {
		t.Errorf("expected %d occurrences of 'Channel:', got %d", expectChannelCount, strings.Count(trim.String(result), "Channel:"))
	}

	if !strings.Contains(trim.String(result), "Channel: 100") {
		t.Errorf("expected 'Channel: 100', got %s", result)
	}
}

// TestResolucaoNaPraticaExercicio7 tests the function ResolucaoNaPraticaExercicio7
// by capturing its output and verifying the number of occurrences of the string "Channel:"
// in the output. It expects 101 occurrences of "Channel:" and checks if the output contains
// the specific string "99 	 9". If the expectations are not met, the test will fail with
// an appropriate error message.
func TestResolucaoNaPraticaExercicio7(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio7)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expectChannelCount := 101

	trim := trim.New()

	if strings.Contains(trim.String(result), strconv.Itoa(expectChannelCount)) {
		t.Errorf("expected 99 occurrences of 'Channel:', got %d", strings.Count(trim.String(result), "Channel:"))
	}

	if !strings.Contains(trim.String(result), "99 	 9") {
		t.Errorf("expected 'Channel: 100', got %s", result)
	}
}
