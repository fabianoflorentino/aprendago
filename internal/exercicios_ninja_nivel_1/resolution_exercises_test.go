// Package exercicios_ninja_nivel_1 provides solutions to various exercises
// and a questionnaire related to basic Go programming concepts. It includes
// functions that demonstrate variable declarations, type conversions, and
// formatted printing. Additionally, it contains a function to handle a
// questionnaire about Go programming fundamentals.
package exercicios_ninja_nivel_1

import (
	"fmt"
	"testing"
)

// TestResolucaoNaPraticaExercicio1 demonstrates a simple exercise where it declares three variables
// with different types (int, string, and bool), formats them into a single string using fmt.Sprintf,
// and then passes the formatted string to the FormatResolucaoExercicios function for further processing.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)
	expected := "42 James Bond true"

	if resolucao != expected {
		t.Errorf("Expected %v, got %v", expected, resolucao)
	}

	// Additional checks for individual variable values
	if x != 42 {
		t.Errorf("Expected x to be 42, got %v", x)
	}

	if y != "James Bond" {
		t.Errorf("Expected y to be 'James Bond', got %v", y)
	}

	if z != true {
		t.Errorf("Expected z to be true, got %v", z)
	}
}

// TestResolucaoNaPraticaExercicio2 demonstrates the declaration of variables with zero values in Go.
// It declares three variables: an integer `x`, a string `y`, and a boolean `z`.
// These variables are then formatted into a single string using `fmt.Sprintf`,
// which is passed to the `FormatResolucaoExercicios` function for further processing.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	var x int
	var y string
	var z bool

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)
	expected := "0  false"

	if resolucao != expected {
		t.Errorf("Expected %v, got %v", expected, resolucao)
	}

	// Additional checks for individual variable zero values
	if x != 0 {
		t.Errorf("Expected x to be 0, got %v", x)
	}

	if y != "" {
		t.Errorf("Expected y to be an empty string, got %v", y)
	}

	if z != false {
		t.Errorf("Expected z to be false, got %v", z)
	}
}

// TestResolucaoNaPraticaExercicio3 demonstrates the creation and formatting of variables in Go.
// It initializes three variables: an integer `x` with the value 42, a string `y` with the value "James Bond",
// and a boolean `z` with the value true. These variables are then formatted into a single string using
// `fmt.Sprintf` and passed to the `FormatResolucaoExercicios` function for further processing.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)
	expected := "42 James Bond true"

	if resolucao != expected {
		t.Errorf("Expected %v, got %v", expected, resolucao)
	}

	// Additional checks for individual variable values
	if x != 42 {
		t.Errorf("Expected x to be 42, got %v", x)
	}

	if y != "James Bond" {
		t.Errorf("Expected y to be 'James Bond', got %v", y)
	}

	if z != true {
		t.Errorf("Expected z to be true, got %v", z)
	}
}

// TestResolucaoNaPraticaExercicio4 demonstrates the creation of a custom type `ninja`,
// assigns a value to a variable of this type, formats it as a string, and prints
// the type of the variable. The formatted string is passed to a function for further
// processing.
//
// The function performs the following steps:
// 1. Defines a new type `ninja` as an alias for `int`.
// 2. Creates a variable `x` of type `ninja` and assigns it the value 42.
// 3. Formats the value of `x` as a string using `fmt.Sprintf`.
// 4. Passes the formatted string to `format.FormatResolucaoExercicios` for further processing.
// 5. Prints the type of the variable `x` using `fmt.Printf`.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	type ninja int

	x := ninja(42)

	resolucao := fmt.Sprintf("%v", x)
	expected := "42"

	if resolucao != expected {
		t.Errorf("Expected %v, got %v", expected, resolucao)
	}

	// Additional checks for type and value
	if fmt.Sprintf("%T", x) != "exercicios_ninja_nivel_1.ninja" {
		t.Errorf("Expected type exercicios_ninja_nivel_1.ninja, got %T", x)
	}

	if x != 42 {
		t.Errorf("Expected value 42, got %v", x)
	}
}

// TestResolucaoNaPraticaExercicio5 demonstrates type conversion in Go.
// It defines a new type 'ninja' as an alias for int, assigns a value to a variable of this new type,
// converts it to a standard int type, and then formats the types of both variables into a string.
// Finally, it calls a function to format and display the result.
func TestResolucaoNaPraticaExercicio5(t *testing.T) {
	type ninja int

	var x ninja
	var y int

	x = 42
	y = int(x)

	resolucao := fmt.Sprintf("%T\n%T", x, y)
	expected := "exercicios_ninja_nivel_1.ninja\nint"

	if resolucao != expected {
		t.Errorf("Expected %v, got %v", expected, resolucao)
	}

	// Additional checks for type and value
	if fmt.Sprintf("%T", x) != "exercicios_ninja_nivel_1.ninja" {
		t.Errorf("Expected type exercicios_ninja_nivel_1.ninja, got %T", x)
	}

	if fmt.Sprintf("%T", y) != "int" {
		t.Errorf("Expected type int, got %T", y)
	}
}
