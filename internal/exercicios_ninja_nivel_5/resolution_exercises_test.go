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
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
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

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected result.
// The expected output is a formatted string containing names and their favorite ice cream flavors.
// If the captured output does not match the expected output, the test will fail and report the difference.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result := output.Capture(ResolucaoNaPraticaExercicio2)

	expect := `
de Tal
	Chocolate
	Morango
	Baunilha
da Silva
	Pistache
	Creme
	Coco
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected result. The test
// checks if the output contains the expected formatted strings for a truck
// and a sedan, including their attributes such as the number of doors and color.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result := output.Capture(ResolucaoNaPraticaExercicio3)

	expect := `
Caminhonete: {veiculo:{portas:4 cor:Preto} tracaoNasQuatro:true}
Sedan: {veiculo:{portas:2 cor:Branco} modeloLuxo:true}
Portas da caminhonete: 4
Cor do sedan: Branco
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
