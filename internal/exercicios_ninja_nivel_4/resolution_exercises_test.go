// Package exercicios_ninja_nivel_4 contains solutions and tests for exercises
// aimed at practicing and improving Go programming skills. The exercises cover
// various topics and concepts, providing hands-on experience and reinforcing
// understanding of the language.
package exercicios_ninja_nivel_4

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	var output string

	func() {
		defer func() {
			if rr := recover(); rr != nil {
				output = fmt.Sprint(rr)
			}
		}()

		output = captureOutput(ResolucaoNaPraticaExercicio1)
	}()

	expect := `
Resolução:
1, 2, 3, 4, 5
`
	if !strings.Contains(strings.TrimSpace(output), strings.TrimSpace(expect)) {
		t.Errorf(expectTemplate, expect, output)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected output. It uses a
// deferred function to recover from any panic that might occur during the execution
// of ResolucaoNaPraticaExercicio2 and captures the panic message if it happens.
// The test checks if the captured output contains the expected string, ignoring
// leading and trailing whitespace, and reports an error if it does not match.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	var output string

	func() {
		defer func() {
			if r := recover(); r != nil {
				output = fmt.Sprint(r)
			}
		}()
		output = captureOutput(ResolucaoNaPraticaExercicio2)
	}()

	expect := `
Resolução:

Slice: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
Tipo: []string
`

	if !strings.Contains(strings.TrimSpace(output), strings.TrimSpace(expect)) {
		t.Errorf(expectTemplate, expect, output)
	}
}

// captureOutput captures and returns the output of a function that writes to the standard output.
// It temporarily redirects os.Stdout to a pipe, executes the provided function, and then restores os.Stdout.
// The captured output is returned as a string.
func captureOutput(resolutionFunction func()) string {
	read, write, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = write

	resolutionFunction()

	write.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, read)
	return buf.String()
}
