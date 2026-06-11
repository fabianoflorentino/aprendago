package exercicios_ninja_nivel_11

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/output"
)

func TestSqrtHappyPath_Coverage(t *testing.T) {
	result, err := sqrt(25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("expected 42, got %v", result)
	}
}

func TestResolucaoNaPraticaExercicio1_Coverage(t *testing.T) {
	out := output.New()
	result, err := out.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		t.Fatalf("unexpected error capturing output: %v", err)
	}

	if !strings.Contains(result, "James") {
		t.Errorf("expected output to contain 'James', got: %s", result)
	}
}

func TestResolucaoNaPraticaExercicio2_Coverage(t *testing.T) {
	out := output.New()
	result, err := out.Capture(ResolucaoNaPraticaExercicio2)
	if err != nil {
		t.Fatalf("unexpected error capturing output: %v", err)
	}

	if !strings.Contains(result, "James") {
		t.Errorf("expected output to contain 'James', got: %s", result)
	}
}
