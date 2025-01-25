package exercicios_ninja_nivel_8

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
)

// const (
// 	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
// )

func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Exemplo 1

User 1: {"First":"M","Age":54}
User 2: {"First":"Q","Age":64}
User 3: {"First":"Moneypenny","Age":27}
User 4: {"First":"James","Age":32}

Exemplo 2

[{"First":"M","Age":54},{"First":"Q","Age":64},{"First":"Moneypenny","Age":27},{"First":"James","Age":32}]`

	// Utility function to normalize strings
	normalize := func(input string) string {
		lines := strings.Split(input, "\n")
		for i, line := range lines {
			lines[i] = strings.TrimSpace(line) // Remove leading/trailing spaces
		}
		trimmed := strings.Join(lines, "\n") // Rejoin lines
		return strings.TrimSpace(trimmed)    // Trim leading/trailing newlines
	}

	normalizedResult := normalize(result)
	normalizedExpect := normalize(expect)

	// Log outputs for debugging
	// t.Logf("Normalized Expected: %q", normalizedExpect)
	// t.Logf("Normalized Result: %q", normalizedResult)

	if normalizedResult != normalizedExpect {
		t.Errorf("Expected output does not match result.\nWant:\n%q\nGot:\n%q", normalizedExpect, normalizedResult)
	}
}
