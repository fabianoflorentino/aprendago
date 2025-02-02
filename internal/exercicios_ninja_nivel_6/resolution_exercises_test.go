// Package exercicios_ninja_nivel_6 contains test functions for exercises in level 6 of the Ninja Go course.
// These tests are designed to validate the solutions to the practical exercises provided in the course.
package exercicios_ninja_nivel_6

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

// expectTemplate is a format string used to display the expected and actual values
// in a test output. It includes placeholders for the expected value ("want") and
// the actual value ("got"), each followed by a newline for better readability.
const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

// TestResolucaoNaPraticaExercicio1 tests the function ResolucaoNaPraticaExercicio1
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output and
// the trim.String method to normalize the strings before comparison.
// If the captured output does not contain the expected result, the test fails
// and an error message is logged.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Inteiro: 42
Inteiro&Texto: 42 e Olá, mundo!
	`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected result.
// It verifies that the function correctly calculates and prints the sum of a variadic
// argument and a slice. If the captured output does not match the expected output,
// the test will fail and log an error.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Soma do variádico: 55
Soma do slice: 66
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output and
// the strings.Contains method to check if the captured output matches the expected output.
// If the output does not match, it logs an error with the expected and actual results.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Execução do defer ocorre ao final do contexto ao qual ela pertence.
Execução do defer ocorreu ao final do contexto.
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4
// by capturing its output and comparing it to the expected result.
// It logs an error if the captured output does not match the expected output.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio4)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Nome: Fabiano Florentino
Idade: 39
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio5 tests the function ResolucaoNaPraticaExercicio5
// by capturing its output and comparing it to the expected output.
// It verifies that the function correctly calculates and prints the area of a square and a circle.
func TestResolucaoNaPraticaExercicio5(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio5)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Área do quadrado: 100.000000
Área do círculo: 31.400000
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio6 tests the function ResolucaoNaPraticaExercicio6
// by capturing its output and comparing it to the expected output.
// It uses the output.Capture method to capture the function's output and
// the trim.String method to normalize the strings before comparison.
// If the captured output does not contain the expected string, the test fails.
func TestResolucaoNaPraticaExercicio6(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio6)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Anonymous function executed.
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio7 tests the function ResolucaoNaPraticaExercicio7
// by capturing its output and comparing it to the expected result.
// It logs an error if the captured output does not match the expected output.
func TestResolucaoNaPraticaExercicio7(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio7)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Função atribuída a uma variável.
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio8 tests the function ResolucaoNaPraticaExercicio8
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output
// and the strings.Contains method to check if the captured output
// contains the expected string. If the output does not match the expected
// result, it logs an error.
func TestResolucaoNaPraticaExercicio8(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio8)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Função retornada.
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio9 tests the function ResolucaoNaPraticaExercicio9
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output
// and the trim.String method to normalize the strings before comparison.
// If the captured output does not contain the expected string, the test fails.
func TestResolucaoNaPraticaExercicio9(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio9)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Função passada como argumento
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio10 tests the function ResolucaoNaPraticaExercicio10
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture method to capture the function's output and
// the strings.Contains method to check if the captured output contains the expected string.
// If the captured output does not match the expected result, it logs an error.
func TestResolucaoNaPraticaExercicio10(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio10)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
10
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio11 tests the function ResolucaoNaPraticaExercicio11
// by capturing its output and comparing it to the expected output. If the captured
// output does not match the expected output, the test fails and logs an error.
func TestResolucaoNaPraticaExercicio11(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio11)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
- Uma das melhores maneiras de aprender é ensinando.
- Para este exercício escolha o seu código favorito dentre os que vimos estudando funções. Pode ser das aulas ou da seção de exercícios. Então:
		- Faça download e instale isso aqui: https://obsproject.com/
		- Grave um vídeo onde *você* ensina o tópico em questão
		- Faça upload do vídeo no YouTube
		- Compartilhe o vídeo no Twitter e me marque no tweet (@ellenkorbes)
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
