// Package exercicios_ninja_nivel_4 contains solutions and tests for exercises
// aimed at practicing and improving Go programming skills. The exercises cover
// various topics and concepts, providing hands-on experience and reinforcing
// understanding of the language.
package exercicios_ninja_nivel_4

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
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
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
1, 2, 3, 4, 5
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the function ResolucaoNaPraticaExercicio2
// by capturing its output and comparing it to the expected output. It uses a
// deferred function to recover from any panic that might occur during the execution
// of ResolucaoNaPraticaExercicio2 and captures the panic message if it happens.
// The test checks if the captured output contains the expected string, ignoring
// leading and trailing whitespace, and reports an error if it does not match.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:

Slice: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
Tipo: []string
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the function ResolucaoNaPraticaExercicio3
// by capturing its output and comparing it to the expected result. It uses the
// output.Capture utility to capture the function's output and the trim.New utility
// to normalize the strings before comparison. If the captured output does not
// contain the expected result, the test fails with an error message.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:

Slice: 1, 2, 3, 4
Tipo: []string
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the function ResolucaoNaPraticaExercicio4
// by capturing its output and comparing it to the expected result. The test
// verifies that the function correctly appends elements to slices and produces
// the expected output. If the actual output does not match the expected output,
// the test will fail and report the discrepancy.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio4)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
append52: [42 43 44 45 46 47 48 49 50 51 52]
append53to55: [42 43 44 45 46 47 48 49 50 51 52 53 54 55]
appendSliceY: [42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio5 tests the function ResolucaoNaPraticaExercicio5
// by capturing its output and comparing it to the expected result.
// It uses the output.Capture utility to capture the function's output
// and the trim.New utility to normalize the strings before comparison.
// If the captured output does not contain the expected result, the test fails
// and an error message is displayed.
func TestResolucaoNaPraticaExercicio5(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio5)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
[42 43 44 48 49 50 51]
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio6 tests the function ResolucaoNaPraticaExercicio6.
// It captures the output of the function and compares it with the expected result.
// If the captured output does not match the expected result, the test fails with an error message.
func TestResolucaoNaPraticaExercicio6(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio6)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
len: 26 cap: 26
Estados: Acre, Alagoas, Amapá, Amazonas, Bahia, Ceará, Espírito Santo, Goiás, Maranhão, Mato Grosso, Mato Grosso do Sul, Minas Gerais, Pará, Paraíba, Paraná, Pernambuco, Piauí, Rio de Janeiro, Rio Grande do Norte, Rio Grande do Sul, Rondônia, Roraima, Santa Catarina, São Paulo, Sergipe, Tocantins
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio7 tests the function ResolucaoNaPraticaExercicio7
// by capturing its output and comparing it to the expected result. The test
// verifies that the output contains the expected formatted strings for names,
// surnames, and favorite hobbies.
func TestResolucaoNaPraticaExercicio7(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio7)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
Nome: Fabiano, Sobrenome: Florentino, Hobby favorito: Programar
Nome: Fulano, Sobrenome: de Tal, Hobby favorito: Jogar bola
Nome: Ciclano, Sobrenome: da Silva, Hobby favorito: Assistir filmes
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio8 tests the function ResolucaoNaPraticaExercicio8
// by capturing its output and comparing it with the expected result. The test
// checks if the output contains the expected formatted string with names and hobbies.
// If the output does not match the expected result, the test will fail and report an error.
func TestResolucaoNaPraticaExercicio8(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio8)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
Sobrenome_Nome: florentino_fabiano
  Hobby 1: Programar
  Hobby 2: Jogar bola
  Hobby 3: Assistir filmes
Sobrenome_Nome: de_tal_fulano
  Hobby 1: Programar
  Hobby 2: Jogar bola
  Hobby 3: Assistir filmes
Sobrenome_Nome: da_silva_ciclano
  Hobby 1: Programar 2
  Hobby 2: Jogar bola 2
  Hobby 3: Assistir filmes 2
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio9 tests the function ResolucaoNaPraticaExercicio9
// by capturing its output and comparing it to the expected result. The test
// verifies that the output contains the expected formatted strings for different
// individuals and their hobbies. If the output does not match the expected result,
// the test will fail and report the discrepancy.
func TestResolucaoNaPraticaExercicio9(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio9)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
Sobrenome_Nome: florentino_fabiano
  Hobby 1: Programar
  Hobby 2: Jogar bola
  Hobby 3: Assistir filmes
Sobrenome_Nome: de_tal_fulano
  Hobby 1: Programar
  Hobby 2: Jogar bola
  Hobby 3: Assistir filmes
Sobrenome_Nome: da_silva_ciclano
  Hobby 1: Programar 2
  Hobby 2: Jogar bola 2
  Hobby 3: Assistir filmes 2
Sobrenome_Nome: de_tal_ciclano
  Hobby 1: Programar 3
  Hobby 2: Jogar bola 3
  Hobby 3: Assistir filmes 3
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}

// TestResolucaoNaPraticaExercicio10 tests the function ResolucaoNaPraticaExercicio10.
// It captures the output of the function and compares it with the expected output.
// If the captured output does not match the expected output, the test fails and an error is reported.
func TestResolucaoNaPraticaExercicio10(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio10)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Resolução:
Sobrenome_Nome: florentino_fabio
  Hobby 1: Programar
  Hobby 2: Jogar bola
  Hobby 3: Assistir filmes
Sobrenome_Nome: de_tal_fulano
  Hobby 1: Programar
  Hobby 2: Jogar bola
  Hobby 3: Assistir filmes
Sobrenome_Nome: de_tal_ciclano
  Hobby 1: Programar 3
  Hobby 2: Jogar bola 3
  Hobby 3: Assistir filmes 3
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
