package agrupamento_de_dados

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/section"
)

func TestMenu(t *testing.T) {
	// Create a new section with a root directory.
	_, err := section.New(rootDir)
	if err != nil {
		t.Errorf("Error creating section: %v", err)
	}

	// Create a slice of MenuOptions.
	menuOptions := Menu(
		[]string{
			"--array",
			"--slice-literal-composta",
			"--slice-for-range",
			"--slice-fatiando-ou-deletando-de-uma-fatia",
			"--slice-anexando-a-uma-slice",
			"--slice-make",
			"--slice-multi-dimensional",
			"--slice-a-surpresa-do-array-subjacente",
			"--maps-introducao",
			"--maps-range-e-deletando",
		})

	// Check if the menuOptions slice has the expected length.
	if len(menuOptions) != 10 {
		t.Errorf("Expected menuOptions length of 10, got %v", len(menuOptions))
	}

	// Check if the menuOptions slice has the expected options.
	expectedOptions := []string{
		"--array",
		"--slice-literal-composta",
		"--slice-for-range",
		"--slice-fatiando-ou-deletando-de-uma-fatia",
		"--slice-anexando-a-uma-slice",
		"--slice-make",
		"--slice-multi-dimensional",
		"--slice-a-surpresa-do-array-subjacente",
		"--maps-introducao",
		"--maps-range-e-deletando",
	}

	for i, option := range menuOptions {
		if option.Options != expectedOptions[i] {
			t.Errorf("Expected option %v to be %v, got %v", i, expectedOptions[i], option.Options)
		}
	}

	// Check if the menuOptions slice has the expected execution functions.
	for i, option := range menuOptions {
		switch i {
		case 0:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 1:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 2:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 3:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 4:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 5:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 6:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 7:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 8:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		case 9:
			if option.ExecFunc == nil {
				t.Errorf("Expected option %v to have a non-nil execution function", i)
			}
		}
	}
}
