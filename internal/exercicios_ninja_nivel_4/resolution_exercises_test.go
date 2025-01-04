// Package exercicios_ninja_nivel_4 contains solutions and tests for exercises
// aimed at practicing and improving Go programming skills. The exercises cover
// various topics and concepts, providing hands-on experience and reinforcing
// understanding of the language.
package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"
	"testing"
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
	var resolution string
	array := [5]int{1, 2, 3, 4, 5}

	for _, v := range array {
		if v == array[len(array)-1] {
			resolution += fmt.Sprintf("%d", v)
			continue
		}
		resolution += fmt.Sprintf("%d, ", v)
	}
	resolution = strings.TrimSpace(resolution)

	expect := "1, 2, 3, 4, 5"

	if resolution != expect {
		t.Errorf(expectTemplate, expect, resolution)
	}
}
