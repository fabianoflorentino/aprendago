// Package exercicios_ninja_nivel_5 contains exercises and their resolutions
// for the Ninja Level 5 of the Go programming language course. These exercises
// are designed to help learners practice and improve their Go programming skills.
package exercicios_ninja_nivel_5

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

const (
	expectTemplate = "\nwant:\n%s\n, \ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and comparing it to the expected result.
// The expected output is a formatted string with names and their favorite ice cream flavors.
// If the captured output does not match the expected output, the test will fail.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result := output.Capture(ResolucaoNaPraticaExercicio1)

	expect := `
Fulano de Tal
	Chocolate
	Morango
	Baunilha
Ciclano da Silva
	Pistache
	Creme
	Coco
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
