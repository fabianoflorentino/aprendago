// TestResolucaoNaPraticaExercicio1 tests the JSON marshaling of a person struct.
// It creates a person instance, marshals it to JSON, and compares the result with the expected JSON string.
// If the marshaling fails or the result does not match the expected string, the test fails with an appropriate error message.

// TestResolucaoNaPraticaExercicio2 tests the toJSON function for converting a person struct to JSON.
// It creates a person instance, converts it to JSON using the toJSON function, and compares the result with the expected JSON string.
// If the conversion fails or the result does not match the expected string, the test fails with an appropriate error message.
package exercicios_ninja_nivel_11

import (
	"encoding/json"
	"testing"
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
