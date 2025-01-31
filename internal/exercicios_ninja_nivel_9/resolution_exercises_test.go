// Package exercicios_ninja_nivel_9 contains tests and exercises for level 9 of the Ninja Go course.
// This package includes various functions and tests to practice advanced Go programming concepts.
package exercicios_ninja_nivel_9

import (
	"strconv"
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used to display the expected and actual values
// in a test output. It helps in comparing the expected result with the actual result
// by formatting them in a readable way.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and checking if the number of occurrences of the word "Goroutine"
// matches the expected count. If the output does not match the expected count, the test fails.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := 10

	if strings.Count(result, "Goroutine") != expect {
		t.Errorf(expectTemplate, strconv.Itoa(expect), result)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected string.
// It fails the test if the output does not contain the expected string.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Olá, meu nome é fabiano
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3 by capturing its output,
// extracting the expected integer value from the output, and verifying if the value is greater than 50.
// If the value is less than or equal to 50, the test fails with an error message.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := strings.Split(result, ":")[1]
	trim := trim.New()

	expectInt, err := strconv.Atoi(trim.String(expect))
	if err != nil {
		t.Fatalf("failed to convert expect to int: %v", err)
	}

	println(expectInt)

	if expectInt <= 50 {
		t.Errorf(expectTemplate, strconv.Itoa(50), strings.Split(result, ":")[1])
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4
// by capturing its output and verifying that it contains the expected string
// "All goroutines completed". It also checks if the integer value extracted from
// the output is greater than 42000. If any of these conditions are not met, the
// test will fail with an appropriate error message.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio4)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	trim := trim.New()

	expectString := `All goroutines completed`

	if !strings.Contains(trim.String(result), trim.String(expectString)) {
		t.Errorf(expectTemplate, expectString, result)
	}

	expectInt := strings.Split(result, ":")[1]
	expectInt = strings.TrimSpace(expectInt)

	expect, err := strconv.Atoi(expectInt)
	if err != nil {
		t.Fatalf("failed to convert expect to int: %v", err)
	}

	if expect <= 42000 {
		t.Errorf(expectTemplate, strconv.Itoa(40000), expectInt)
	}
}
