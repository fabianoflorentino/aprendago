package exercicios_ninja_nivel_7

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used to display the expected and actual values
// in test output. It helps in comparing the expected result with the actual result
// by formatting them in a readable way.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and logging any errors that occur during the capture process.
// It also prints a message indicating that it will show the memory address of a string variable.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	println("Show the memory address of a string variable")
	println("N/A", result)
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected result.
// It logs an error if the captured output does not match the expected output.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Original -> Nome: Fulano de Tal, Idade: 25
Alterado -> Nome: João Silva, Idade: 30
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
