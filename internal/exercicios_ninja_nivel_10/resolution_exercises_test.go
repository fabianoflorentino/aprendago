package exercicios_ninja_nivel_10

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
		t.Fatalf("unexpected error: %v", err)
	}

	expect := `
Value: 42
`
	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
