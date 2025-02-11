// TestResolucaoNaPraticaExercicio1 tests the JSON marshaling of a person struct.
// It creates a person instance, marshals it to JSON, and compares the result with the expected JSON string.
// If the marshaling fails or the result does not match the expected string, the test fails with an appropriate error message.

// TestResolucaoNaPraticaExercicio2 tests the toJSON function for converting a person struct to JSON.
// It creates a person instance, converts it to JSON using the toJSON function, and compares the result with the expected JSON string.
// If the conversion fails or the result does not match the expected string, the test fails with an appropriate error message.
package exercicios_ninja_nivel_11

import (
	"encoding/json"
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

// TestResolucaoNaPraticaExercicio1 tests the JSON marshaling of a 'person' struct.
// It creates a 'person' instance, marshals it to JSON, and compares the result
// with the expected JSON string. If there is an error during marshaling or the
// marshaled JSON does not match the expected string, the test will fail.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `{"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]}`
	if string(bs) != expected {
		t.Errorf("Expected %v, got %v", expected, string(bs))
	}
}

// TestResolucaoNaPraticaExercicio2 tests the toJSON function to ensure it correctly
// converts a person struct to its JSON representation. It creates a person instance,
// converts it to JSON, and compares the result with the expected JSON string. If the
// conversion fails or the result does not match the expected string, the test fails.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `{"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]}`

	if string(bs) != expected {
		t.Errorf("Expected %v, got %v", expected, string(bs))
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected output.
// It uses the output.Capture method to capture the function's output and
// the strings.Contains method to check if the captured output contains the expected string.
// If the output does not match the expected string, the test fails with an error message.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `
erro especial
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expected)) {
		t.Errorf(expectTemplate, expected, result)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4
// by capturing its output and checking if it contains the expected substring.
// It ensures that the function handles the specific error message correctly.
// If the expected substring is not found in the output, the test will fail.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio4)
	if err != nil {
		t.Fatalf("Unexpected error capturing output: %v", err)
	}

	expectedSubstring := "math error: 50.2289 N 99.4656 W número negativo: -10.23"

	// Verifica apenas a parte relevante da mensagem de erro.
	if !strings.Contains(result, expectedSubstring) {
		t.Errorf("Expected substring not found.\nWant: %q\nGot: %q", expectedSubstring, result)
	}
}
