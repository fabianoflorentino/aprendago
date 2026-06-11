package exercicios_ninja_nivel_12

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/output"
)

func TestResolucaoNaPraticaExercicio2_Coverage(t *testing.T) {
	out := output.New()
	result, err := out.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		t.Fatalf("unexpected error capturing output: %v", err)
	}

	if !strings.Contains(result, "Não há resolução") {
		t.Errorf("expected output to contain 'Não há resolução', got: %s", result)
	}
}

func TestResolucaoNaPraticaExercicio3_Coverage(t *testing.T) {
	out := output.New()
	result, err := out.Capture(ResolucaoNaPraticaExercicio3)
	if err != nil {
		t.Fatalf("unexpected error capturing output: %v", err)
	}

	if !strings.Contains(result, "Não há resolução") {
		t.Errorf("expected output to contain 'Não há resolução', got: %s", result)
	}
}
