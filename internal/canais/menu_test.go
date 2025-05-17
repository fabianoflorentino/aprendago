package canais

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/section"
)

func TestMenuCanais(t *testing.T) {
	// Create a new section with a root directory.
	_, err := section.New(rootDir)
	if err != nil {
		t.Errorf("Error creating section: %v", err)
	}

	// Create a slice of MenuOptions for MenuCanais.
	menuOptions := Menu(
		[]string{
			flagEntendendoCanais,
			flagCanaisDirecionaisUtilizandoCanais,
			flagRangeEClose,
			flagSelectStatement,
			flagCommaOkExpression,
			flagConvergencia,
			flagConvergenciaExample,
			flagConvergenciaExampleChanString,
			flagDivergencia,
			flagDivergenciaExample,
			flagDivergenciaExampleWithFunc,
			flagContext,
		})

	// Check if the menuOptions slice has the expected length.
	if len(menuOptions) != 12 {
		t.Errorf("Expected menuOptions length of 12, got %v", len(menuOptions))
	}

	// Check if the menuOptions slice has the expected options.
	expectedOptions := []string{
		"--entendendo-canais",
		"--canais-direcionais-utilizando-canais",
		"--range-e-close",
		"--select",
		"--a-expressao-comma-ok",
		"--convergencia",
		"--convergencia --example",
		"--convergencia --example --chan-string",
		"--divergencia",
		"--divergencia --example",
		"--divergencia --example --with-func",
		"--context",
	}

	for i, option := range menuOptions {
		if option.Options != expectedOptions[i] {
			t.Errorf("Expected option %v to be %v, got %v", i, expectedOptions[i], option.Options)
		}
	}

	// Check if the menuOptions slice has the expected execution functions.
	for i, option := range menuOptions {
		if option.ExecFunc == nil {
			t.Errorf("Expected option %v to have a non-nil execution function", i)
		}
	}
}
